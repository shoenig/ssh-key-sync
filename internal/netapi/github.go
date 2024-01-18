package netapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/hashicorp/go-set/v2"
	"github.com/shoenig/ignore"
	"github.com/shoenig/ssh-key-sync/internal/config"
	"github.com/shoenig/ssh-key-sync/internal/logs"
	"github.com/shoenig/ssh-key-sync/internal/ssh"
)

const (
	githubKeysPath          = "/users/USERNAME/keys"
	githubAcceptHeaderValue = "application/vnd.github.v3+json"
)

// NewGithubClient creates a Client that can be used to communicate
// with the github API.
func NewGithubClient(args config.Arguments) Client {
	return &githubClient{
		url:    args.GitHubAPI,
		client: httpClient,
		logger: logs.New(args.Verbose),
	}
}

type githubClient struct {
	url    string
	client *http.Client
	logger *log.Logger
}

type githubKeySSH struct {
	ID  int    `json:"id"`
	Key string `json:"key"`
}

func (g *githubClient) GetKeys(user string) (*set.Set[ssh.Key], error) {
	url := appendToURL(g.url, strings.Replace(githubKeysPath, "USERNAME", user, 1))

	var jsonKeys []githubKeySSH

	if err := g.doGet(url, &jsonKeys); err != nil {
		return nil, err
	}

	keys := set.New[ssh.Key](len(jsonKeys))
	for _, jsonKey := range jsonKeys {
		parsed, err := ssh.ParseKey(jsonKey.Key, true)
		if err != nil {
			g.logger.Printf("failed to parse key: %v", err)
			continue
		}
		keys.Insert(parsed)
	}
	return keys, nil
}

func (g *githubClient) doGet(url string, i interface{}) error {
	g.logger.Printf("acquire github keys from %q", url)

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	request.Header.Set("Accept", githubAcceptHeaderValue)
	request.Header.Set("User-Agent", userAgent)

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
