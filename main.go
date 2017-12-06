// Author hoenig

package main

import (
	"fmt"
	"os"

	"github.com/shoenig/ssh-key-sync/internal/command"
	"github.com/shoenig/ssh-key-sync/internal/config"
	"github.com/shoenig/ssh-key-sync/internal/github"
)

func main() {
	args := config.ParseArguments()
	loader := config.NewLoader(args.ConfigFile)
	client := github.NewClient(nil)

	execer := command.NewExecer(loader, client)
	if err := execer.Exec(); err != nil {
		fmt.Printf("ssh-key-sync had error: %s\n", err)
		os.Exit(1)
	}
}
