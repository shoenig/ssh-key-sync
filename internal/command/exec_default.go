//go:build !linux

package command

import (
	"github.com/shoenig/go-landlock"
)

func lockdown(keyFile string) error {
	return nil
}
