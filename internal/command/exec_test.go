// Author hoenig
// License MIT

package command

import (
	"testing"

	"github.com/shoenig/ssh-key-sync/internal/config"
	"github.com/shoenig/ssh-key-sync/internal/config/configtest"
	"github.com/shoenig/ssh-key-sync/internal/github/githubtest"
	"github.com/shoenig/ssh-key-sync/internal/ssh"
	"github.com/shoenig/ssh-key-sync/internal/ssh/sshtest"
	"github.com/stretchr/testify/require"
)

func Test_Exec(t *testing.T) {
	loader := &configtest.Loader{}
	client := &githubtest.Client{}
	reader := &sshtest.KeysReader{}

	loader.On("Load").Return(config.Options{
		System: []config.User{
			{User: "bob", AuthorizedKeysFile: "/tmp/home/bob/keys.txt"},
			{User: "sally", AuthorizedKeysFile: "/tmp/home/sally/.ssh/authorized_keys"},
		},
		Github: config.Github{
			URL: "https://api.github.com",
			Accounts: []config.GithubAccount{
				{Username: "billybob", SystemUser: "bob"},
				{Username: "sadsally", SystemUser: "sally"},
			},
		},
	}, nil).Once()

	bob1 := ssh.Key{Managed: false, Value: "aaaaaaa", User: "bob", Host: "b1"}
	bob2 := ssh.Key{Managed: false, Value: "bbbbbbb", User: "bob", Host: "b2"}
	bob3 := ssh.Key{Managed: true, Value: "ccccccc", User: "bob", Host: "b1"}
	bob4 := ssh.Key{Managed: true, Value: "ddddddd", User: "bob", Host: "b3"}
	sally1 := ssh.Key{Managed: false, Value: "jjjjjjj", User: "sally", Host: "s1"}
	sally2 := ssh.Key{Managed: false, Value: "kkkkkkk", User: "sally", Host: "s2"}
	sally3 := ssh.Key{Managed: true, Value: "lllllll", User: "sally", Host: "s3"}

	reader.On("ReadKeys", "/tmp/home/bob/keys.txt").Return(
		[]ssh.Key{bob1, bob2}, nil,
	)

	reader.On("ReadKeys", "/tmp/home/sally/.ssh/authorized_keys").Return(
		[]ssh.Key{sally1, sally2}, nil,
	)

	client.On("GetKeys", "billybob").Return(
		[]ssh.Key{bob3, bob4}, nil,
	).Once()

	client.On("GetKeys", "sadsally").Return(
		[]ssh.Key{sally3}, nil,
	).Once()

	ex := NewExecer(loader, reader, client)
	ex.(*execer).fakeChown = true

	err := ex.Exec()
	require.NoError(t, err)

	loader.AssertExpectations(t)
	reader.AssertExpectations(t)
	client.AssertExpectations(t)
}
