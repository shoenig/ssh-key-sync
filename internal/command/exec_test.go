// Author hoenig

package command

import (
	"testing"

	"github.com/shoenig/ssh-key-sync/internal/config"
	"github.com/shoenig/ssh-key-sync/internal/config/configtest"
	"github.com/shoenig/ssh-key-sync/internal/github/githubtest"
	"github.com/stretchr/testify/require"
)

func Test_Exec(t *testing.T) {
	loader := &configtest.Loader{}
	client := &githubtest.Client{}

	loader.On("Load").Return(config.Options{
		Github: config.Github{
			URL: "https://api.github.com",
			Accounts: []config.GithubAccount{
				{Username: "billybob", AuthorizedKeysFile: "/home/bob/keys.txt"},
				{Username: "sadsally", AuthorizedKeysFile: "/home/sally/.ssh/authorized_keys"},
			},
		},
	}, nil).Once()

	client.On("GetKeys", "billybob").Return(
		[]string{"one", "two"}, nil,
	).Once()

	client.On("GetKeys", "sadsally").Return(
		[]string{"key1"}, nil,
	).Once()

	execer := NewExecer(loader, client)

	err := execer.Exec()
	require.NoError(t, err)

	loader.AssertExpectations(t)
	client.AssertExpectations(t)
}
