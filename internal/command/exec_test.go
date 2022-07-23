package command

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/shoenig/test/must"
	"github.com/shoenig/ssh-key-sync/internal/config"
	"github.com/shoenig/ssh-key-sync/internal/netapi"
	"github.com/shoenig/ssh-key-sync/internal/ssh"
)

func mkdirs(t *testing.T, fp string) {
	dirs := filepath.Dir(fp)
	err := os.MkdirAll(dirs, 0700)
	must.NoError(t, err)
}

func Test_Exec(t *testing.T) {
	keyBob := "/tmp/home/bob/keys.txt"
	keySally := "/tmp/home/sally/.ssh/authorized_keys"
	keyNed := "/tmp/home/ned/.ssh/authorized_keys"

	mkdirs(t, keyBob)
	mkdirs(t, keySally)
	mkdirs(t, keyNed)

	githubClient := netapi.NewClientMock(t)
	gitlabClient := netapi.NewClientMock(t)
	reader := ssh.NewKeysReaderMock(t)

	defer githubClient.MinimockFinish()
	defer gitlabClient.MinimockFinish()
	defer reader.MinimockFinish()

	opts := &config.Options{
		System: []config.User{
			{User: "bob", AuthorizedKeysFile: keyBob},
			{User: "sally", AuthorizedKeysFile: keySally},
			{User: "ned", AuthorizedKeysFile: keyNed},
		},
		Github: config.Github{
			URL: "https://api.github.com",
			Accounts: []config.WebAccount{
				{Username: "billybob", SystemUser: "bob"},
				{Username: "sadsally", SystemUser: "sally"},
			},
		},
		Gitlab: config.Gitlab{
			URL:   "https://gitlab.local",
			Token: "abcdefg",
			Accounts: []config.WebAccount{
				{Username: "bobbo", SystemUser: "bob"},
				{Username: "ned", SystemUser: "ned"},
			},
		},
	}

	bob1 := ssh.Key{Managed: false, Value: "aaaaaaa", User: "bob", Host: "b1"}
	bob2 := ssh.Key{Managed: false, Value: "bbbbbbb", User: "bob", Host: "b2"}
	bob3 := ssh.Key{Managed: true, Value: "ccccccc", User: "bob", Host: "b1"}
	bob4 := ssh.Key{Managed: true, Value: "ddddddd", User: "bob", Host: "b3"}
	bob5 := ssh.Key{Managed: true, Value: "eeeeeee", User: "bob", Host: "b5"}
	sally1 := ssh.Key{Managed: false, Value: "jjjjjjj", User: "sally", Host: "s1"}
	sally2 := ssh.Key{Managed: false, Value: "kkkkkkk", User: "sally", Host: "s2"}
	sally3 := ssh.Key{Managed: true, Value: "lllllll", User: "sally", Host: "s3"}
	ned1 := ssh.Key{Managed: true, Value: "ppppppp", User: "ned", Host: "n1"}

	reader.ReadKeysMock.When(keyBob).Then([]ssh.Key{bob1, bob2}, nil)
	reader.ReadKeysMock.When(keySally).Then([]ssh.Key{sally1, sally2}, nil)
	reader.ReadKeysMock.When(keyNed).Then([]ssh.Key{}, nil)
	githubClient.GetKeysMock.When("billybob").Then([]ssh.Key{bob3, bob4}, nil)
	githubClient.GetKeysMock.When("sadsally").Then([]ssh.Key{sally3}, nil)
	gitlabClient.GetKeysMock.When("bobbo").Then([]ssh.Key{bob5}, nil)
	gitlabClient.GetKeysMock.When("ned").Then([]ssh.Key{ned1}, nil)

	ex := NewExecer(reader, githubClient, gitlabClient)
	ex.(*execer).fakeChown = true

	err := ex.Exec(opts)
	must.NoError(t, err)
}
