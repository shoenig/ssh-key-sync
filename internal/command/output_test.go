package command

import (
	"os"
	"strings"
	"testing"
	"time"

	"github.com/shoenig/ssh-key-sync/internal/ssh"
	"github.com/shoenig/test/must"
)

func TestOutput_writeToFile(t *testing.T) {
	f, err := os.CreateTemp("", "")
	must.NoError(t, err)

	err = f.Close()
	must.NoError(t, err)

	err = writeToFile(f.Name(), "hello")
	must.NoError(t, err)

	b, err := os.ReadFile(f.Name())
	must.NoError(t, err)
	must.Eq(t, "hello", string(b))
}

func TestOutput_generateFileContent(t *testing.T) {
	now := time.Date(2022, 1, 2, 3, 4, 0, 0, time.UTC)
	keys := []ssh.Key{
		{
			Managed: false,
			Value:   "abc123",
			User:    "alice",
			Host:    "h1",
		},
		{
			Managed: true,
			Value:   "def234",
			User:    "bob",
			Host:    "h2",
		},
		{
			Managed: true,
			Value:   "ghi345",
			User:    "carla",
			Host:    "h3",
		},
	}
	s := generateFileContent(keys, now)

	exp, err := os.ReadFile("../../hack/tests/generatedOutput1.txt")
	must.NoError(t, err)
	must.Eq(t, strings.TrimSpace(s), strings.TrimSpace(string(exp)))
}
