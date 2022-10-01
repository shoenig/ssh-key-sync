package netapi

import (
	"fmt"
	"net/http"
	"time"

	"github.com/hashicorp/go-set"
	"github.com/shoenig/ssh-key-sync/internal/ssh"
)

// A Client is used to acquire keys from an API service like
// github/gitlab (public or internal).
type Client interface {
	GetKeys(user string) (*set.Set[ssh.Key], error)
}

var (
	// a shared http client with a default timeout
	httpClient = &http.Client{Timeout: 10 * time.Second}

	// the user-agent to use for all http requests
	userAgent = fmt.Sprintf("ssh-key-sync bot/v2 (https://github.com/shoenig/ssh-key-sync)")
)
