//go:build linux

package command

import (
	"github.com/shoenig/go-landlock"
)

func paths(keyFile string) []*landlock.Path {
	return []*landlock.Path{
		landlock.Certs(),
		landlock.DNS(),
		landlock.File(keyFile, "rw"),
	}
}
