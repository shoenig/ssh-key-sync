package command

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"sort"
	"strconv"
	"time"

	"github.com/shoenig/ssh-key-sync/internal/config"
	"github.com/shoenig/ssh-key-sync/internal/logs"
	"github.com/shoenig/ssh-key-sync/internal/netapi"
	"github.com/shoenig/ssh-key-sync/internal/ssh"
)

type Exec interface {
	Execute(config.Arguments) error
}

func NewExec(
	verbose bool,
	reader ssh.KeysReader,
	githubClient netapi.Client,
) Exec {
	return &exec{
		logger:       logs.New(verbose),
		reader:       reader,
		githubClient: githubClient,
	}
}

type exec struct {
	logger       *log.Logger
	reader       ssh.KeysReader
	githubClient netapi.Client
}

func (e *exec) Execute(args config.Arguments) error {
	return e.processUser(args.SystemUser, args.GitHubUser, args.AuthorizedKeys)
}

func (e *exec) processUser(systemUser, githubUser, keyFile string) error {
	e.logger.Printf("process local user %s from %s@github", systemUser, githubUser)

	// 1) ensure the authorized key file exists, and belongs to user
	if err := e.touch(keyFile, systemUser); err != nil {
		return fmt.Errorf("failed to touch %q for user %q: %w", keyFile, systemUser, err)
	}

	// 2) load existing keys from authorization file
	localKeys, err := e.reader.ReadKeys(keyFile)
	if err != nil {
		return fmt.Errorf("failed to load keys from %q for user %q: %w", keyFile, systemUser, err)
	}
	e.logger.Printf("loaded %d existing keys for user %q", len(localKeys), systemUser)

	// 3) maybe load keys from github account
	githubKeys, err := e.getKeys(e.githubClient, githubUser)
	if err != nil {
		return fmt.Errorf("failed to fetch keys from github user %q: %w", githubUser, err)
	}
	e.logger.Printf("retrieved %d keys for github user: %s", len(githubKeys), githubUser)

	// 5) combine the keys, purging old managed keys with the new set
	newKeys := combine(onlyUnmanaged(localKeys), githubKeys)
	content := generateFileContent(newKeys, time.Now())

	// 6) write the new file content to the authorized keys file
	return e.writeToFile(keyFile, systemUser, content)
}

func (e *exec) getKeys(client netapi.Client, githubUser string) ([]ssh.Key, error) {
	return client.GetKeys(githubUser)
}

func combine(keySets ...[]ssh.Key) []ssh.Key {
	result := make([]ssh.Key, 0, 10)
	for _, keySet := range keySets {
		result = append(result, keySet...)
	}
	sort.Sort(ssh.KeySorter(result))
	return result
}

func onlyUnmanaged(keys []ssh.Key) []ssh.Key {
	unmanaged := make([]ssh.Key, 0, len(keys))
	for _, key := range keys {
		if !key.Managed {
			unmanaged = append(unmanaged, key)
		}
	}
	return unmanaged
}

func (e *exec) touch(path, username string) error {
	e.logger.Printf("touch key file for %s: %s", username, path)
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_RDONLY, 0600)
	if err != nil {
		return err
	}

	if err := f.Sync(); err != nil {
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}

	u, err := user.Lookup(username)
	if err != nil {
		return err
	}

	uid, err := strconv.Atoi(u.Uid)
	if err != nil {
		return err
	}

	gid, err := strconv.Atoi(u.Gid)
	if err != nil {
		return err
	}

	return os.Chown(path, uid, gid)
}
