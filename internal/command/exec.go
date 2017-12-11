// Author hoenig

package command

import (
	"fmt"

	"github.com/shoenig/ssh-key-sync/internal/config"
	"github.com/shoenig/ssh-key-sync/internal/github"
)

type Execer interface {
	Exec() error
}

func NewExecer(
	loader config.Loader,
	client github.Client,
) Execer {
	return &execer{
		loader: loader,
		client: client,
	}
}

type execer struct {
	loader config.Loader
	client github.Client
}

func (e *execer) Exec() error {
	opts, err := e.loader.Load()
	if err != nil {
		return fmt.Errorf("failed to load config: %v", err)
	}

	for _, account := range opts.Github.Accounts {
		user := account.Username
		githubKeys, err := e.client.GetKeys(user)
		if err != nil {
			return fmt.Errorf("failed to retrieve github keys for user %q: %v", user, err)
		}

		fmt.Printf("retrieved %d github keys for user %q: %v\n", len(githubKeys), user, githubKeys)
	}

	return nil
}
