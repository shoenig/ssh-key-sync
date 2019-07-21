package netapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"

	"gophers.dev/cmds/ssh-key-sync/internal/ssh"

	"gophers.dev/pkgs/ignore"
)

const (
	githubURL               = "https://api.github.com"
	githubKeysPath          = "/users/USERNAME/keys"
	githubAcceptHeaderValue = "application/vnd.github.v3+json"
)

// NewGithubClient creates a Client that can be used to communicate
// with the github API.
func NewGithubClient(options Optioner) Client {
	opts := options.Options()
	if opts == nil {
		opts = &Options{}
	}
	return &githubClient{
		url:    opts.url(githubURL),
		client: httpClient,
	}
}

type githubClient struct {
	url    string
	client *http.Client
}

type githubKeySSH struct {
	ID  int    `json:"id"`
	Key string `json:"key"`
}

func (g *githubClient) GetKeys(user string) ([]ssh.Key, error) {
	url := appendToURL(g.url, strings.Replace(githubKeysPath, "USERNAME", user, 1))

	var jsonkeys []githubKeySSH

	if err := g.doGet(url, &jsonkeys); err != nil {
		return nil, err
	}

	keys := make([]ssh.Key, 0, len(jsonkeys))
	for _, jsonkey := range jsonkeys {
		parsed, err := ssh.ParseKey(jsonkey.Key, true)
		if err != nil {
			return nil, err
		}
		keys = append(keys, parsed)
	}

	sort.Sort(ssh.KeySorter(keys))

	return keys, nil
}

func (g *githubClient) doGet(url string, i interface{}) error {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	request.Header.Set("Accept", githubAcceptHeaderValue)
	request.Header.Set("User-Agent", useragent)

	response, err := g.client.Do(request)
	if err != nil {
		return err
	}
	defer ignore.Drain(response.Body)

	if response.StatusCode >= 400 {
		return fmt.Errorf("request to %q returned code %d", url, response.StatusCode)
	}

	return json.NewDecoder(response.Body).Decode(i)
}

func appendToURL(url, path string) string {
	trimmedURL := strings.TrimSuffix(url, "/")
	trimmedPath := strings.TrimPrefix(path, "/")
	return trimmedURL + "/" + trimmedPath
}
