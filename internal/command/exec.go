// Author hoenig

package command

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
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
		loader: loader,
		reader: reader,
		client: client,
	}
}

type execer struct {
	loader config.Loader
	reader ssh.KeysReader
	client github.Client
}

func (e *execer) Exec() error {
	opts, err := e.loader.Load()
	if err != nil {
		return fmt.Errorf("failed to load config: %v", err)
	}

	// pull this out once we do more than just github
	for _, account := range opts.Github.Accounts {
		if err := e.processGithub(account); err != nil {
			return err
		}
	}

	return nil
}

func (e *execer) processGithub(account config.GithubAccount) error {
	// 1) ensure the authorized keyfile exists
	if err := touch(account.AuthorizedKeysFile); err != nil {
		return err
	}

	// 2) read existing keys from authorized keys file
	localKeys, err := e.reader.ReadKeys(account.AuthorizedKeysFile)
	if err != nil {
		return err
	}

	fmt.Printf("read %d local keys from %s\n", len(localKeys), account.AuthorizedKeysFile)

	// 3) retrieve managed keys from github
	user := account.Username
	githubKeys, err := e.client.GetKeys(user)
	if err != nil {
		return fmt.Errorf("failed to retrieve github keys for user %q: %v", user, err)
	}

	fmt.Printf("retrieved %d github keys for user %q: %v\n", len(githubKeys), user, githubKeys)

	// 4) combine the keys for the complete new set
	newKeys := combine(onlyUnmanaged(localKeys), githubKeys)

	s := generateFileContent(newKeys, time.Now())
	fmt.Println("new keys:\n", s)

	return nil
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

func touch(path string) error {
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

	return f.Close()
}
