package command

import (
	"os"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/go-set"
	"github.com/shoenig/ssh-key-sync/internal/logs"
	"github.com/shoenig/ssh-key-sync/internal/ssh"
	"github.com/shoenig/test/must"
	"oss.indeed.com/go/libtime/libtimetest"
)

type mockKeysReader struct {
}

func (r *mockKeysReader) ReadKeys(filename string) (*set.Set[ssh.Key], error) {
	return set.From[ssh.Key]([]ssh.Key{{
		Managed: false,
		Value:   "abc123",
		User:    "alice",
		Host:    "a1",
	}, {
		Managed: true,
		Value:   "def345",
		User:    "bob",
		Host:    "a2",
	}}), nil
}

type mockClient struct {
}

func (c *mockClient) GetKeys(username string) (*set.Set[ssh.Key], error) {
	return set.From[ssh.Key]([]ssh.Key{{
		Managed: true,
		Value:   "def345",
		User:    "bob",
		Host:    "a2",
	}, {
		Managed: true,
		Value:   "ghi456",
		User:    "carla",
		Host:    "c3",
	}}), nil
}

func TestExec_Execute(t *testing.T) {
	exp, err := os.ReadFile("../../hack/tests/generatedOutput2.txt")
	must.NoError(t, err)

	writer := func(filename, content string) error {
		must.Eq(t, "/nothing/keys", filename)
		must.Eq(t, strings.TrimSpace(string(exp)), strings.TrimSpace(content))
		return nil
	}

	clock := libtimetest.NewClockMock(t)
	clock.NowMock.Return(time.Date(2022, 10, 2, 8, 53, 0, 0, time.UTC))

	e := &exec{
		logger:       logs.New(true),
		reader:       new(mockKeysReader),
		githubClient: new(mockClient),
		clock:        clock,
		writeKeyFile: writer,
	}

	err = e.processUser("bob", "bob-gh", "/nothing/keys")
	must.NoError(t, err)
}
