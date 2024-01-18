package command

import (
	"os"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/go-set/v2"
	"github.com/shoenig/ssh-key-sync/internal/logs"
	"github.com/shoenig/ssh-key-sync/internal/ssh"
	"github.com/shoenig/test/must"
	"oss.indeed.com/go/libtime/libtimetest"
)

type mockKeysReader struct {
	readKeysResult *set.Set[ssh.Key]
}

func (r *mockKeysReader) ReadKeys(_ string) (*set.Set[ssh.Key], error) {
	if r.readKeysResult != nil {
		return r.readKeysResult, nil
	}

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
	getKeysResult *set.Set[ssh.Key]
}

func (c *mockClient) GetKeys(_ string) (*set.Set[ssh.Key], error) {
	if c.getKeysResult != nil {
		return c.getKeysResult, nil
	}
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
		prune:        false,
		logger:       logs.New(true),
		reader:       new(mockKeysReader),
		githubClient: new(mockClient),
		clock:        clock,
		writeKeyFile: writer,
	}

	err = e.processUser("bob", "bob-gh", "/nothing/keys")
	must.NoError(t, err)
}

func TestExec_Execute_prune(t *testing.T) {
	exp, err := os.ReadFile("../../hack/tests/generatedOutput3.txt")
	must.NoError(t, err)

	writer := func(filename, content string) error {
		must.Eq(t, "/nothing/keys", filename)
		must.Eq(t, strings.TrimSpace(string(exp)), strings.TrimSpace(content))
		return nil
	}

	clock := libtimetest.NewClockMock(t)
	clock.NowMock.Return(time.Date(2022, 10, 2, 8, 53, 0, 0, time.UTC))

	e := &exec{
		prune:        true,
		logger:       logs.New(true),
		reader:       new(mockKeysReader),
		githubClient: new(mockClient),
		clock:        clock,
		writeKeyFile: writer,
	}

	err = e.processUser("bob", "bob-gh", "/nothing/keys")
	must.NoError(t, err)
}

func TestExec_Execute_prune_empty(t *testing.T) {
	writer := func(filename, _ string) error {
		must.Eq(t, "/nothing/keys", filename)
		return nil
	}

	clock := libtimetest.NewClockMock(t)
	clock.NowMock.Return(time.Date(2022, 10, 2, 8, 53, 0, 0, time.UTC))

	mkr := &mockKeysReader{
		readKeysResult: set.From([]ssh.Key{{
			Managed: false,
			Value:   "abc123",
			User:    "alice",
			Host:    "a1",
		}}),
	}

	mc := &mockClient{
		getKeysResult: set.New[ssh.Key](0),
	}

	e := &exec{
		prune:        true,
		logger:       logs.New(true),
		reader:       mkr,
		githubClient: mc,
		clock:        clock,
		writeKeyFile: writer,
	}

	err := e.processUser("bob", "bob-gh", "/nothing/keys")
	must.ErrorContains(t, err, "no keys!")
}

func TestExec_combine(t *testing.T) {
	locals := set.From[ssh.Key]([]ssh.Key{{
		Managed: false,
		Value:   "abc123",
		User:    "alice",
		Host:    "a1",
	}, {
		Managed: false,
		Value:   "def345",
		User:    "bob",
		Host:    "a2",
	}})

	github := set.From[ssh.Key]([]ssh.Key{{
		Managed: true,
		Value:   "yyy111",
		User:    "alice",
		Host:    "a1",
	}, {
		Managed: true,
		Value:   "zzz222",
		User:    "bob",
		Host:    "a2",
	}})

	t.Run("normal", func(t *testing.T) {
		e := &exec{prune: false}
		result := e.combine(locals, github)
		must.Len(t, 4, result)
	})

	t.Run("prune", func(t *testing.T) {
		e := &exec{prune: true, logger: logs.New(true)}
		result := e.combine(locals, github)
		must.Len(t, 2, result)
	})
}
