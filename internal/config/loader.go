// Author hoenig

package config

import (
	"encoding/json"
	"os"
)

//go:generate mockery -interface=Loader -package=configtest

type Loader interface {
	Load() (Options, error)
}

type Github struct {
	URL      string          `json:"url"`
	Accounts []GithubAccount `json:"accounts"`
}

type GithubAccount struct {
	Username           string `json:"username"`
	AuthorizedKeysFile string `json:"authorized_keys_file"`
}

type Options struct {
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
