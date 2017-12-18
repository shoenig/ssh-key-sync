// Author hoenig
// License MIT

package main

import (
	"fmt"
	"os"

	"github.com/shoenig/ssh-key-sync/internal/command"
	"github.com/shoenig/ssh-key-sync/internal/config"
	"github.com/shoenig/ssh-key-sync/internal/github"
	"github.com/shoenig/ssh-key-sync/internal/ssh"
)

func main() {
	args := config.ParseArguments()
	loader := config.NewLoader(args.ConfigFile)
	reader := ssh.NewKeysReader()
	client := github.NewClient(nil)

	execer := command.NewExecer(loader, reader, client)
	if err := execer.Exec(); err != nil {
		fmt.Printf("ssh-key-sync had error: %s\n", err)
		os.Exit(1)
	}
}
