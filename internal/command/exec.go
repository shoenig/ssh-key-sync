package command

import (
	"fmt"
	"log"
	"path/filepath"
	"sort"

	"github.com/hashicorp/go-set/v2"
	"github.com/shoenig/go-landlock"
	"github.com/shoenig/ssh-key-sync/internal/config"
	"github.com/shoenig/ssh-key-sync/internal/logs"
	"github.com/shoenig/ssh-key-sync/internal/netapi"
	"github.com/shoenig/ssh-key-sync/internal/ssh"
	"oss.indeed.com/go/libtime"
)

type Exec interface {
	Execute(config.Arguments) error
}

func NewExec(
	prune bool,
	verbose bool,
	reader ssh.KeysReader,
	githubClient netapi.Client,
) Exec {
	return &exec{
		prune:        prune,
		logger:       logs.New(verbose),
		reader:       reader,
		githubClient: githubClient,
		clock:        libtime.SystemClock(),
		writeKeyFile: writeToFile,
	}
}

type exec struct {
	prune        bool
	logger       *log.Logger
	reader       ssh.KeysReader
	githubClient netapi.Client
	clock        libtime.Clock
	writeKeyFile func(filename, content string) error
}

func (e *exec) Execute(args config.Arguments) error {
	switch args.AuthorizedKeys {
	case "":
		args.AuthorizedKeys = filepath.Join("/home", args.SystemUser, ".ssh", "authorized_keys")
		e.logger.Printf("using default output authorized_keys file (%s)", args.AuthorizedKeys)
	default:
		e.logger.Printf("using configured output authorized_keys file (%s)", args.AuthorizedKeys)
	}

	if err := lockdown(args.AuthorizedKeys); err != nil {
		return err
	}

	if err := e.processUser(args.SystemUser, args.GitHubUser, args.AuthorizedKeys); err != nil {
		return err
	}

	return nil
}

func (e *exec) processUser(systemUser, githubUser, keyFile string) error {
	e.logger.Printf("process local user %s from %s@github", systemUser, githubUser)

	// 1) load existing keys from authorization file
	localKeys, err := e.reader.ReadKeys(keyFile)
	if err != nil {
		return fmt.Errorf("failed to load keys from %q for user %q: %w", keyFile, systemUser, err)
	}
	e.logger.Printf("loaded %d existing keys for user %q", localKeys.Size(), systemUser)

	// 2) maybe load keys from github account
	githubKeys, err := e.githubClient.GetKeys(githubUser)
	if err != nil {
		return fmt.Errorf("failed to fetch keys from github user %q: %w", githubUser, err)
	}
	e.logger.Printf("retrieved %d keys for github user: %s", githubKeys.Size(), githubUser)

	// 3) combine the keys, purging old managed keys with the new set
	newKeys := e.combine(localKeys, githubKeys)
	if len(newKeys) == 0 {
		return fmt.Errorf("no keys! refusing to write empty set of keys")
	}

	// 4) write the new file content to the authorized keys file
	content := generateFileContent(newKeys, e.clock.Now())
	return e.writeKeyFile(keyFile, content)
}

func (e *exec) combine(local, gh *set.Set[ssh.Key]) []ssh.Key {
	if e.prune {
		e.logger.Printf("pruning %d non-github managed keys", local.Size())
		result := gh.Slice()
		sort.Sort(ssh.KeySorter(result))
		return result
	}
	union := local.Union(gh)
	result := union.Slice()
	sort.Sort(ssh.KeySorter(result))
	return result
}

func lockdown(keyfile string) error {
	if landlock.Available() {
		ll := landlock.New(paths(keyfile)...)
		return ll.Lock(landlock.OnlySupported)
	}
	return nil
}
