package netapi

import (
	"net/http"
	"strings"
	"testing"

	"github.com/shoenig/test/must"
)

const (
	githubKeysResponse = `
[
  {
    "id": 15746919,
    "key": "ssh-rsa AAAAB3Nzaeyij"
  },
  {
    "id": 16608270,
    "key": "ssh-rsa AAAAB3NzaZ1yk="
  },
  {
    "id": 20879474,
    "key": "ssh-rsa AAAAB3NzaC1yc2E"
  }
]
`
)

func Test_GithubClient_GetKeys(t *testing.T) {
	opts, ts := makeServer(func(w http.ResponseWriter, r *http.Request) {
		suffix := strings.HasSuffix(r.URL.Path, "/users/bobby/keys")
		must.True(t, suffix)

		_, _ = w.Write([]byte(githubKeysResponse))
	})
	t.Cleanup(ts.Close)

	client := NewGithubClient(opts)

	keys, err := client.GetKeys("bobby")
	must.NoError(t, err)

	// sorted by ascii
	must.LenSlice(t, 3, keys)
	must.EqCmp(t, "ssh-rsa AAAAB3Nzaeyij", keys[2].Value)
	must.EqCmp(t, "ssh-rsa AAAAB3NzaZ1yk=", keys[1].Value)
	must.EqCmp(t, "ssh-rsa AAAAB3NzaC1yc2E", keys[0].Value)
}
