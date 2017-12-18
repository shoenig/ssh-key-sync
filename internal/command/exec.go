// Author hoenig

package command

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"sort"
	"strconv"
	"time"

	"github.com/shoenig/ssh-key-sync/internal/config"
	"github.com/shoenig/ssh-key-sync/internal/github"
	"github.com/shoenig/ssh-key-sync/internal/ssh"
)

type Execer interface {
	Exec() error
}

func NewExecer(
	loader config.Loader,
	reader ssh.KeysReader,
	client github.Client,
) Execer {
	return &execer{
		loader:    loader,
		reader:    reader,
		client:    client,
		fakeChown: false,
	}
}

type execer struct {
	loader    config.Loader
	reader    ssh.KeysReader
	client    github.Client
	fakeChown bool
}

func (e *execer) Exec() error {
	opts, err := e.loader.Load()
	if err != nil {
		return fmt.Errorf("failed to load config: %v", err)
	}

	users2keyfiles := opts.SystemUsers()
	users2github := opts.GithubUsers()

	for username, keyfile := range users2keyfiles {
		if err := e.processUser(username, keyfile, users2github); err != nil {
			return err
		}
	}

	return nil
}
func (e *execer) processUser(user, keyfile string, users2github map[string]string) error {
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

	// 3) load keys from github account
	githubKeys, err := e.keysFromGithub(user, users2github)
	if err != nil {
		return fmt.Errorf("failed to fetch keys from github for user %q: %v", user, err)
	}
	fmt.Printf("retrieved %d keys for user %q from github\n", len(githubKeys), user)

	// 4) combine the keys, purging old managed keys with the new set
	newKeys := combine(onlyUnmanaged(localKeys), githubKeys)
	content := generateFileContent(newKeys, time.Now())

	// 5) write the new file content to the authorized keys file
	return writeToFile(keyfile, content)
}

func (e *execer) keysFromGithub(user string, users2github map[string]string) ([]ssh.Key, error) {
	githubUsername, exists := users2github[user]
	if !exists {
		return nil, nil
	}
	return e.client.GetKeys(githubUsername)
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
	dirs := filepath.Dir(path)

	if err := os.MkdirAll(dirs, 0700); err != nil {
		return err
	}

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
