//go:build linux

package command

import (
	"github.com/shoenig/go-landlock"
)

func lockdown(keyFile string) error {
	ll := landlock.New(
		landlock.Certs(),
		landlock.DNS(),
		landlock.File(keyFile, "rw"),
	)
	return ll.Lock(landlock.Mandatory)
}
