package meta

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Version(t *testing.T) {
	versionRe := regexp.MustCompile(`^[\d]+.[\d]+.[\d]+$`)
	matches := versionRe.MatchString(Version)
	require.True(t, matches)
}
