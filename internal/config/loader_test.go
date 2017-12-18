// Author hoenig

package config

import "testing"

func Test_Loader(t *testing.T) {
	l := NewLoader("../../hack/tests/config.1")
	opts, err := l.Load()
	if err != nil {
		t.Fatal(err)
	}

	accounts := []struct {
		Username           string
		AuthorizedKeysFile string
	}{
		{Username: "alice", AuthorizedKeysFile: "/tmp/home/alice/authorized_keys"},
		{Username: "bob", AuthorizedKeysFile: "/tmp/home/bob/authorized_keys"},
	}

	if len(accounts) != len(opts.Github.Accounts) {
		t.Fatalf(
			"number of expected accounts (%d) and number of actual accounts (%d) differ",
			len(accounts),
			len(opts.Github.Accounts),
		)
	}

	for i := range accounts {
		result := opts.Github.Accounts[i]
		exp := accounts[i]
		if result.Username != exp.Username {
			t.Fatalf("account[%d].Username does not match, got: %v, exp: %v", i, result, exp)
		}
		if result.AuthorizedKeysFile != exp.AuthorizedKeysFile {
			t.Fatalf("account[%d].AuthorizedKeysFile does not match, got: %v, exp: %v", i, result, exp)
		}
	}
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
