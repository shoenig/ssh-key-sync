// Author hoenig

package config

import (
	"encoding/json"
	"os"
)

type Options struct {
	Github []struct {
		Username           string `json:"username"`
		AuthorizedKeysFile string `json:"authorized_keys_file"`
	} `json:"github"`
}

type Loader interface {
	Load() (Options, error)
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
