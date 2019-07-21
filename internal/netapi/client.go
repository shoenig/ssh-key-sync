package netapi

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"gophers.dev/cmds/ssh-key-sync/internal/meta"
	"gophers.dev/cmds/ssh-key-sync/internal/ssh"
)

//go:generate go run github.com/gojuno/minimock/cmd/minimock -g -i Client -s _mock.go

// A Client is used to acquire keys from an API service like
// github/gitlab (public or internal).
type Client interface {
	GetKeys(user string) ([]ssh.Key, error)
}

// An Optioner returns some Options.
type Optioner interface {
	Options() *Options
}

// Options represents configuration parameters available
// for reaching API services like github and gitlab.
type Options struct {
	URL   string `json:"url"`
	Token string `json:"token"`
}

func (o Options) url(defaultURL string) string {
	if o.URL == "" {
		return transportize(defaultURL)
	}
	return transportize(o.URL)
}

func transportize(url string) string {
	if !strings.HasPrefix(url, "http") {
		return "https://" + url
	}
	return url
}

var (
	// a shared http client with a default timeout
	httpClient = &http.Client{Timeout: 10 * time.Second}

	// the user-agent to use for all http requests
	useragent = fmt.Sprintf("ssh-key-sync bot/%s (https://gophers.dev/cmds/ssh-key-sync)", meta.Version)
)
