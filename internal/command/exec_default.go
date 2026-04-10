//go:build !linux

package command

import (
	"github.com/shoenig/go-landlock"
)

func paths(_ string) []*landlock.Path {
	return nil
}
