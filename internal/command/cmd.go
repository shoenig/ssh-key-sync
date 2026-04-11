package command

import (
	"fmt"
	"os"

	"cattlecloud.net/go/babycli"
	"github.com/shoenig/ssh-key-sync/internal/config"
	"github.com/shoenig/ssh-key-sync/internal/netapi"
	"github.com/shoenig/ssh-key-sync/internal/ssh"
	"github.com/shoenig/ssh-key-sync/version"
)

func Invoke(args []string) babycli.Code {
	return babycli.New(&babycli.Configuration{
		Arguments: args,
		Version:   version.Version,
		Top: &babycli.Component{
			Name:        "ssh-key-sync",
			Description: "Sync SSH public keys from GitHub to authorized_keys",
			Flags: babycli.Flags{
				{Long: "verbose", Short: "v", Type: babycli.BooleanFlag, Help: "print verbose logging"},
				{Long: "prune", Short: "p", Type: babycli.BooleanFlag, Help: "delete all keys not found in GitHub"},
				{Long: "system-user", Short: "u", Type: babycli.StringFlag, Help: "specify the unix system user"},
				{Long: "authorized-keys", Type: babycli.StringFlag, Help: "override the output authorized_keys file"},
				{Long: "github-user", Short: "g", Type: babycli.StringFlag, Require: true, Help: "specify the GitHub user"},
				{Long: "github-api", Type: babycli.StringFlag, Help: "specify the GitHub API endpoint"},
			},
			Function: run,
		},
	}).Run()
}

func run(c *babycli.Component) babycli.Code {
	githubUser := c.GetString("github-user")

	cfg := config.NewArguments(
		c.GetBool("verbose"),
		c.GetBool("prune"),
		c.GetString("system-user"),
		c.GetString("authorized-keys"),
		githubUser,
		c.GetString("github-api"),
	)

	reader := ssh.NewKeysReader()
	githubClient := netapi.NewGithubClient(cfg)
	exe := NewExec(cfg.Prune, cfg.Verbose, reader, githubClient)

	if err := exe.Execute(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "unable to execute: %v", err)
		return babycli.Failure
	}

	return babycli.Success
}
