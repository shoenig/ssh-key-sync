package netapi

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shoenig/test/must"
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
	must.EqCmp(t, "https://code.net", o.url("code.net"))
	must.EqCmp(t, "http://code.net", o.url("http://code.net"))

	o = &Options{
		URL: "api.gitlab.net",
	}
	must.EqCmp(t, "https://api.gitlab.net", o.url("code.net"))
}
