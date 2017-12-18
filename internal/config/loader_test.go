// Author hoenig
// License MIT

package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Loader(t *testing.T) {
	l := NewLoader("../../hack/tests/config.1")
	opts, err := l.Load()
	if err != nil {
		t.Fatal(err)
	}

	exp := Options{
		System: []User{
			{User: "bobby", AuthorizedKeysFile: "/tmp/home/bobby/authorized_keys"},
			{User: "alice", AuthorizedKeysFile: "/tmp/home/alice/authorized_keys"},
		},
		Github: Github{
			URL: "github.com",
			Accounts: []GithubAccount{
				{Username: "alice", SystemUser: "alice"},
				{Username: "bob", SystemUser: "bobby"},
			},
		},
	}

	require.Equal(t, exp, opts)
}

func Test_Loader_noFile(t *testing.T) {
	l := NewLoader("/path/does/not/ever/exist/for/anybody")
	_, err := l.Load()
	if err == nil {
		t.Fatalf("err should not have been nil for nonexistent config file")
	}
}

func Test_Loader_badFormat(t *testing.T) {
	l := NewLoader("../../hack/tests/config.2")
	_, err := l.Load()
	if err == nil {
		t.Fatalf("err should not have been nil for invalid json in config file")
	}
}
