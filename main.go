package main

import (
	"fmt"
	"os"

	"github.com/shoenig/ssh-key-sync/internal/command"
	"github.com/shoenig/ssh-key-sync/internal/config"
	"github.com/shoenig/ssh-key-sync/internal/netapi"
	"github.com/shoenig/ssh-key-sync/internal/ssh"
)

func main() {
	args := config.ParseArguments()
	if args.GitHubUser == "" {
		_, _ = fmt.Fprintf(os.Stderr, "ssh-key-sync requires --github-user\n")
		os.Exit(1)
	}

	reader := ssh.NewKeysReader()
	githubClient := netapi.NewGithubClient(args)

	exec := command.NewExec(args.Verbose, reader, githubClient)
	if err := exec.Execute(args); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ssh-key-sync failed with error: %s\n", err)
		os.Exit(1)
	}
}
