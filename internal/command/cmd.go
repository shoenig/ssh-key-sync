package command

import (
	"errors"

	"github.com/shoenig/ssh-key-sync/internal/config"
	"github.com/shoenig/ssh-key-sync/internal/netapi"
	"github.com/shoenig/ssh-key-sync/internal/ssh"
)

func Start(args []string) error {
	arguments := config.ParseArguments(args[0], args[1:])
	if arguments.GitHubUser == "" {
		arguments.Usage()
		return errors.New("missing required argument(s)")
	}

	reader := ssh.NewKeysReader()
	githubClient := netapi.NewGithubClient(arguments)
	exe := NewExec(arguments.Verbose, reader, githubClient)
	return exe.Execute(arguments)
}
