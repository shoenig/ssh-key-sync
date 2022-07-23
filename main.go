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
	loader := config.NewLoader(args.ConfigFile)
	reader := ssh.NewKeysReader()

	opts, err := loader.Load()
	if err != nil {
		fmt.Printf("ssh-key-sync failed to load config: %v\n", err)
		os.Exit(1)
	}

	githubClient := netapi.NewGithubClient(opts.Github)
	gitlabClient := netapi.NewGitlabClient(opts.Gitlab)

	execer := command.NewExecer(reader, githubClient, gitlabClient)
	if err := execer.Exec(opts); err != nil {
		fmt.Printf("ssh-key-sync had error: %s\n", err)
		os.Exit(1)
	}
}
