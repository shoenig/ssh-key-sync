package config

import (
	"os/user"
	"testing"

	"github.com/shoenig/test/must"
)

const program = "ssh-key-sync"

func TestSystemUser(t *testing.T) {
	t.Run("from argument", func(t *testing.T) {
		args := ParseArguments(program, []string{"--system-user", "cathy"})
		must.Eq(t, "cathy", args.SystemUser)
	})

	t.Run("from environment", func(t *testing.T) {
		t.Setenv("USER", "athena")
		args := ParseArguments(program, nil)
		must.Eq(t, "athena", args.SystemUser)
	})

	t.Run("from process", func(t *testing.T) {
		u, err := user.Current()
		must.NoError(t, err)
		t.Setenv("USER", "")
		args := ParseArguments(program, nil)
		must.Eq(t, u.Username, args.SystemUser)
	})
}

func TestPrune(t *testing.T) {
	t.Run("no", func(t *testing.T) {
		args := ParseArguments(program, []string{})
		must.False(t, args.Prune)
	})

	t.Run("yes", func(t *testing.T) {
		args := ParseArguments(program, []string{"--prune"})
		must.True(t, args.Prune)
	})
}
