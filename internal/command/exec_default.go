//go:build !linux

package command

import (
	"github.com/shoenig/go-landlock"
)

func paths(keyFile string) []*landlock.Path {
	return nil
}
