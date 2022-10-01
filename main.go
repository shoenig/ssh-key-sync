package main

import (
	"os"

	"github.com/shoenig/ssh-key-sync/internal/command"
	"github.com/shoenig/ssh-key-sync/internal/config"
	"github.com/shoenig/ssh-key-sync/internal/logs"
	"github.com/shoenig/ssh-key-sync/internal/netapi"
	"github.com/shoenig/ssh-key-sync/internal/ssh"
)

func main() {
	logger := logs.New(true)
	args := config.ParseArguments(os.Args[0], os.Args[1:])
	if args.GitHubUser == "" {
		logger.Fatal("ssh-key-sync requires --github-user")
	}

	reader := ssh.NewKeysReader()
	githubClient := netapi.NewGithubClient(args)
	exec := command.NewExec(args.Verbose, reader, githubClient)
	if err := exec.Execute(args); err != nil {
		logger.Fatalf("ssh-key-sync failed with error: %s", err)
	}
}
