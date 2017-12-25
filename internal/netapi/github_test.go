// Author hoenig
// License MIT

package netapi

import (
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
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
		if !strings.HasSuffix(r.URL.Path, "/users/bobby/keys") {
			t.Fatal("unexpected path", r.URL.Path)
		}
		w.Write([]byte(githubKeysResponse))
	})
	defer ts.Close()

	client := NewGithubClient(opts)

	keys, err := client.GetKeys("bobby")
	require.NoError(t, err)

	// sorted by ascii
	require.Equal(t, 3, len(keys))
	require.Equal(t, "ssh-rsa AAAAB3Nzaeyij", keys[2].Value)
	require.Equal(t, "ssh-rsa AAAAB3NzaZ1yk=", keys[1].Value)
	require.Equal(t, "ssh-rsa AAAAB3NzaC1yc2E", keys[0].Value)
}
