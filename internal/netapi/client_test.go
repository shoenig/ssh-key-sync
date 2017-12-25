// Author hoenig

package netapi

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Options_url(t *testing.T) {
	o := &Options{
		URL: "",
	}
	require.Equal(t, o.url("code.net"), "https://code.net")
	require.Equal(t, o.url("http://code.net"), "http://code.net")

	o = &Options{
		URL: "api.gitlab.net",
	}
	require.Equal(t, o.url("code.net"), "https://api.gitlab.net")
}
