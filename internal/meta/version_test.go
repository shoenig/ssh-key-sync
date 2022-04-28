package meta

import (
	"regexp"
	"testing"

	"github.com/shoenig/test/must"
)

func Test_Version(t *testing.T) {
	versionRe := regexp.MustCompile(`^\d+.\d+.\d+$`)
	matches := versionRe.MatchString(Version)
	must.True(t, matches)
}
