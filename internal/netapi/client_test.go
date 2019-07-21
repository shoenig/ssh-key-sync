package netapi

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

type optioner struct {
	opts *Options
}

func (o optioner) Options() *Options {
	return o.opts
}

func makeServer(h http.HandlerFunc) (Optioner, *httptest.Server) {
	ts := httptest.NewServer(h)
	opts := &Options{
		URL:   ts.URL,
		Token: "abc123",
	}
	return optioner{opts: opts}, ts
}

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
