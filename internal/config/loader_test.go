package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Loader(t *testing.T) {
	l := NewLoader("../../hack/tests/config.1")
	opts, err := l.Load()
	require.NoError(t, err)

	expOpts := &Options{
		System: []User{
			{User: "bobby", AuthorizedKeysFile: "/tmp/home/bobby/authorized_keys"},
			{User: "alice", AuthorizedKeysFile: "/tmp/home/alice/authorized_keys"},
			{User: "ned", AuthorizedKeysFile: "/tmp/home/ned/authorized_keys"},
		},
		Github: Github{
			URL: "github.com",
			Accounts: []WebAccount{
				{Username: "alice", SystemUser: "alice"},
				{Username: "bob", SystemUser: "bobby"},
			},
		},
		Gitlab: Gitlab{
			URL: "gitlab.com",
			Accounts: []WebAccount{
				{Username: "alison", SystemUser: "alice"},
				{Username: "ned", SystemUser: "ned"},
			},
		},
	}
	require.Equal(t, expOpts, opts)

	expSystemUsers := map[string]string{
		"bobby": "/tmp/home/bobby/authorized_keys",
		"alice": "/tmp/home/alice/authorized_keys",
		"ned":   "/tmp/home/ned/authorized_keys",
	}
	require.Equal(t, expSystemUsers, opts.SystemUsers())

	expGithubUsers := map[string]string{
		"alice": "alice",
		"bobby": "bob",
	}
	require.Equal(t, expGithubUsers, opts.GithubUsers())

	expGitlabUsers := map[string]string{
		"alice": "alison",
		"ned":   "ned",
	}
	require.Equal(t, expGitlabUsers, opts.GitlabUsers())
}

func Test_Loader_noFile(t *testing.T) {
	l := NewLoader("/path/does/not/ever/exist/for/anybody")
	_, err := l.Load()
	require.Error(t, err)
}

func Test_Loader_badFormat(t *testing.T) {
	l := NewLoader("../../hack/tests/config.2")
	_, err := l.Load()
	require.Error(t, err)
}
