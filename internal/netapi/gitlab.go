package netapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"gophers.dev/cmds/ssh-key-sync/internal/ssh"

	"gophers.dev/pkgs/ignore"
)

const (
	gitlabURL            = "https://gitlab.com"
	gitlabUserPath       = "/api/v4/users?username=USERNAME"
	gitlabKeysPath       = "/api/v4/users/USERID/keys"
	gitlabTokenHeaderKey = "Private-Token"
)

func NewGitlabClient(options Optioner) Client {
	opts := options.Options()
	if opts == nil {
		opts = &Options{}
	}

	return &gitlabClient{
		url:    opts.url(gitlabURL),
		token:  opts.Token,
		client: httpClient,
	}
}

type gitlabClient struct {
	url    string
	token  string
	client *http.Client
}

type gitlabUserInfo struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

type gitlabKeySSH struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Key   string `json:"key"`
}

func (g *gitlabClient) GetKeys(user string) ([]ssh.Key, error) {
	// first we need to retrieve the user info for user, which will have their ID
	infoURI := appendToURL(g.url, strings.Replace(gitlabUserPath, "USERNAME", user, 1))

	// the API returns a list, even though only zero or one items will exist
	var userInfos []gitlabUserInfo

	if err := g.doGet(infoURI, &userInfos); err != nil {
		return nil, err
	}

	if len(userInfos) < 1 {
		return nil, errors.Errorf("no gitlab user of username %q", user)
	}

	if len(userInfos) > 1 {
		return nil, errors.Errorf("somehow more than one gitlab user of username %q", user)
	}

	// we have the user's id, now we get their keys
	id := strconv.Itoa(userInfos[0].ID)
	keysURI := appendToURL(g.url, strings.Replace(gitlabKeysPath, "USERID", id, 1))

	var jsonkeys []gitlabKeySSH
	if err := g.doGet(keysURI, &jsonkeys); err != nil {
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

func (g *gitlabClient) doGet(url string, i interface{}) error {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	request.Header.Set(gitlabTokenHeaderKey, g.token)
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
