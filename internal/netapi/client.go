// Author hoenig
// License MIT

package netapi

import (
	"net/http"
	"time"

	"github.com/shoenig/ssh-key-sync/internal/ssh"
)

//go:generate mockery -interface=Client -package netapitest

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
		return defaultURL
	}
	return o.URL
}

var (
	// a shared http client with a default timeout
	httpClient = &http.Client{Timeout: 10 * time.Second}
)
