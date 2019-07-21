package command

import (
	"fmt"
	"os"
	"os/user"
	"sort"
	"strconv"
	"time"

	"gophers.dev/cmds/ssh-key-sync/internal/config"
	"gophers.dev/cmds/ssh-key-sync/internal/netapi"
	"gophers.dev/cmds/ssh-key-sync/internal/ssh"
)

type Execer interface {
	Exec(*config.Options) error
}

func NewExecer(
	reader ssh.KeysReader,
	githubClient netapi.Client,
	gitlabClient netapi.Client,
) Execer {
	return &execer{
		reader:       reader,
		githubClient: githubClient,
		gitlabClient: gitlabClient,

		fakeChown: false,
	}
}

type execer struct {
	reader       ssh.KeysReader
	githubClient netapi.Client
	gitlabClient netapi.Client

	// testing configuration only
	fakeChown bool
}

func (e *execer) Exec(opts *config.Options) error {
	users2keyfiles := opts.SystemUsers()
	users2github := opts.GithubUsers()
	users2gitlab := opts.GitlabUsers()

	for username, keyfile := range users2keyfiles {
		if err := e.processUser(username, keyfile, users2github, users2gitlab); err != nil {
			return err
		}
	}

	return nil
}

func (e *execer) processUser(
	user,
	keyfile string,
	users2github,
	users2gitlab map[string]string,
) error {

	// 1) ensure the authorized key file exists, and belongs to user
	if err := e.touch(keyfile, user); err != nil {
		return fmt.Errorf("failed to touch %q for user %q: %v", keyfile, user, err)
	}

	// 2) load existing keys from authorization file
	localKeys, err := e.reader.ReadKeys(keyfile)
	if err != nil {
		return fmt.Errorf("failed to load keys from %q for user %q: %v", keyfile, user, err)
	}
	fmt.Printf("loaded %d keys for user %q from %q\n", len(localKeys), user, keyfile)

	// 3) maybe load keys from github account
	githubKeys, err := e.getKeys(e.githubClient, user, users2github)
	if err != nil {
		return fmt.Errorf("failed to fetch keys from github for user %q: %v", user, err)
	}
	fmt.Printf("retrieved %d keys for user %q from github\n", len(githubKeys), user)

	// 4) maybe load keys from gitlab account
	gitlabKeys, err := e.getKeys(e.gitlabClient, user, users2gitlab)
	if err != nil {
		return fmt.Errorf("failed to fetch keys from gitlab for user %q: %v", user, err)
	}
	fmt.Printf("retrieved %d keys for user %q from gitlab\n", len(gitlabKeys), user)

	// 5) combine the keys, purging old managed keys with the new set
	newKeys := combine(onlyUnmanaged(localKeys), githubKeys, gitlabKeys)
	content := generateFileContent(newKeys, time.Now())

	// 6) write the new file content to the authorized keys file
	return e.writeToFile(keyfile, user, content)
}

func (e *execer) getKeys(client netapi.Client, user string, system2account map[string]string) ([]ssh.Key, error) {
	username, exists := system2account[user]
	if !exists {
		return nil, nil
	}
	return client.GetKeys(username)
}

func combine(keysets ...[]ssh.Key) []ssh.Key {
	result := make([]ssh.Key, 0, 10)
	for _, keyset := range keysets {
		result = append(result, keyset...)
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

func (e *execer) touch(path, username string) error {
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

	// if we are configured to fake chown, just return success
	if e.fakeChown {
		return nil
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
