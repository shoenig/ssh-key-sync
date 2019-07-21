package netapi

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	gitlabUsersResponse = `
[
  {
    "id": 9422,
    "name": "Alison",
    "username": "alice"
  }
]
`

	gitlabKeysResponse = `
[
  {
    "id": 15746919,
    "title": "alice@a1",
    "key": "ssh-rsa AAAAB3Nzaeyij"
  },
  {
    "id": 16608270,
    "title": "alice@a2",
    "key": "ssh-rsa AAAAB3NzaZ1yk="
  },
  {
    "id": 20879474,
    "title": "alice@a3",
    "key": "ssh-rsa AAAAB3NzaC1yc2E"
  }
]
`
)

func Test_GitlabClient_GetKeys(t *testing.T) {
	opts, ts := makeServer(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/api/v4/users":
			_, _ = w.Write([]byte(gitlabUsersResponse))
		case "/api/v4/users/9422/keys":
			_, _ = w.Write([]byte(gitlabKeysResponse))
		default:
			t.Fatal("unexpected path", r.URL.Path)
		}
	})
	defer ts.Close()

	client := NewGitlabClient(opts)
	keys, err := client.GetKeys("alice")
	require.NoError(t, err)

	require.Equal(t, 3, len(keys))
	require.Equal(t, "ssh-rsa AAAAB3Nzaeyij", keys[2].Value)
	require.Equal(t, "ssh-rsa AAAAB3NzaZ1yk=", keys[1].Value)
	require.Equal(t, "ssh-rsa AAAAB3NzaC1yc2E", keys[0].Value)
}
