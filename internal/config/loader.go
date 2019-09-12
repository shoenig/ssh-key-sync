package config

import (
	"encoding/json"
	"os"

	"gophers.dev/cmds/ssh-key-sync/internal/netapi"
)

//go:generate go run github.com/gojuno/minimock/v3/cmd/minimock -g -i Loader -s _mock.go

type Loader interface {
	Load() (*Options, error)
}

type User struct {
	User               string `json:"user"`
	AuthorizedKeysFile string `json:"authorized_keys_file"`
}

type Github struct {
	URL      string       `json:"url"`
	Accounts []WebAccount `json:"accounts"`
}

func (g Github) Options() *netapi.Options {
	return &netapi.Options{
		URL: g.URL,
	}
}

type WebAccount struct {
	Username   string `json:"username"`
	SystemUser string `json:"system_user"`
}

type Gitlab struct {
	URL      string       `json:"url"`
	Token    string       `json:"token"`
	Accounts []WebAccount `json:"accounts"`
}

func (g Gitlab) Options() *netapi.Options {
	return &netapi.Options{
		URL:   g.URL,
		Token: g.Token,
	}
}

type Options struct {
	System []User `json:"system"`
	Github Github `json:"github"`
	Gitlab Gitlab `json:"gitlab"`
}

func NewLoader(filepath string) Loader {
	return &loader{path: filepath}
}

type loader struct {
	path string
}

func (l *loader) Load() (*Options, error) {
	var opts Options
	f, err := os.Open(l.path)
	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(f).Decode(&opts); err != nil {
		return nil, err
	}

	return &opts, nil
}

// SystemUsers returns a map from local system username to path of
// associated authorized keys file.
func (o Options) SystemUsers() map[string]string {
	user2keyfile := make(map[string]string, len(o.System))
	for _, user := range o.System {
		user2keyfile[user.User] = user.AuthorizedKeysFile
	}
	return user2keyfile
}

// GithubUsers returns a map from local system user to github
// username (for those that have one defined).
func (o Options) GithubUsers() map[string]string {
	user2github := make(map[string]string, len(o.System))
	for _, account := range o.Github.Accounts {
		user2github[account.SystemUser] = account.Username
	}
	return user2github
}

// GitlabUsers returns a map from local system user to gitlab
// username (for those that have one defined).
func (o Options) GitlabUsers() map[string]string {
	user2gitlab := make(map[string]string, len(o.System))
	for _, account := range o.Gitlab.Accounts {
		user2gitlab[account.SystemUser] = account.Username
	}
	return user2gitlab
}
