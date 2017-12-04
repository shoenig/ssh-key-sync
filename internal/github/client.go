// Author hoenig

package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/shoenig/toolkit"
)

const (
	defaultURL  = "https://api.github.com"
	sshKeysPath = "/users/USERNAME/keys"
	apiHeader   = "application/vnd.github.v3+json"
)

// A Client is used to acquire keys from github.com.
type Client interface {
	GetKeys(user string) ([]string, error)
}

type Options struct {
	URL string `json:"url"`
}

func (opts *Options) url() string {
	if opts.URL == "" {
		return defaultURL
	}
	return opts.URL
}

// NewClient creates a Client that can be used to communicate
// with the github API.
func NewClient(opts *Options) Client {
	if opts == nil {
		opts = &Options{}
	}
	return &githubClient{
		url: opts.url(),
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

type githubClient struct {
	url    string
	client *http.Client
}

type sshKey struct {
	ID  int    `json:"id"`
	Key string `json:"key"`
}

func (g *githubClient) GetKeys(user string) ([]string, error) {
	url := combineURL(g.url, strings.Replace(sshKeysPath, "USERNAME", user, 1))

	var sshkeys []sshKey

	if err := g.doGet(url, &sshkeys); err != nil {
		return nil, err
	}

	keys := make([]string, 0, len(sshkeys))
	for _, key := range sshkeys {
		keys = append(keys, key.Key)
	}

	sort.Strings(keys)

	return keys, nil
}

func (g *githubClient) doGet(url string, i interface{}) error {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	request.Header.Set("Accept", apiHeader)

	response, err := g.client.Do(request)
	if err != nil {
		return err
	}
	defer toolkit.Drain(response.Body)

	if response.StatusCode >= 400 {
		return fmt.Errorf("request to %q returned code %d", url, response.StatusCode)
	}

	return json.NewDecoder(response.Body).Decode(i)
}

func combineURL(url, path string) string {
	trimmedURL := strings.TrimSuffix(url, "/")
	trimmedPath := strings.TrimPrefix(path, "/")
	return trimmedURL + "/" + trimmedPath
}
