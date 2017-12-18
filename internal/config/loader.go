// Author hoenig
// License MIT

package config

import (
	"encoding/json"
	"os"
)

//go:generate mockery -interface=Loader -package=configtest

type Loader interface {
	Load() (Options, error)
}

type User struct {
	User               string `json:"user"`
	AuthorizedKeysFile string `json:"authorized_keys_file"`
}

type Github struct {
	URL      string          `json:"url"`
	Accounts []GithubAccount `json:"accounts"`
}

type GithubAccount struct {
	Username   string `json:"username"`
	SystemUser string `json:"system_user"`
}

type Options struct {
	System []User `json:"system"`
	Github Github `json:"github"`
}

func NewLoader(filepath string) Loader {
	return &loader{path: filepath}
}

type loader struct {
	path string
}

func (l *loader) Load() (Options, error) {
	var opts Options
	f, err := os.Open(l.path)
	if err != nil {
		return opts, err
	}

	if err := json.NewDecoder(f).Decode(&opts); err != nil {
		return opts, err
	}

	return opts, nil
}

// SystemUsers returns a map from local system username to path of associated authorized keys file.
func (o Options) SystemUsers() map[string]string {
	user2keyfile := make(map[string]string, len(o.System))
	for _, user := range o.System {
		user2keyfile[user.User] = user.AuthorizedKeysFile
	}
	return user2keyfile
}

// GithubUsers returns a map from local system user to github username (for those that have one defined).
func (o Options) GithubUsers() map[string]string {
	user2github := make(map[string]string, len(o.System))
	for _, account := range o.Github.Accounts {
		user2github[account.SystemUser] = account.Username
	}
	return user2github
}
