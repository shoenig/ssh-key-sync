package config

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

var current string

func init() {
	u, err := user.Current()
	if err != nil {
		panic("unable to lookup current user")
	}
	current = u.Username
}

type Arguments struct {
	Verbose bool
	Prune   bool

	SystemUser     string
	SystemHome     string
	AuthorizedKeys string

	GitHubUser string
	GitHubAPI  string

	usage func()
}

func (a Arguments) Usage() {
	a.usage()
}

func defaultUser() string {
	if u := os.Getenv("USER"); u != "" {
		return u
	}
	return current
}

func ParseArguments(program string, args []string) Arguments {
	flags := flag.NewFlagSet(program, flag.ContinueOnError)
	var arguments Arguments

	flags.BoolVar(
		&arguments.Verbose,
		"verbose", false, "print verbose logging",
	)

	flags.BoolVar(
		&arguments.Prune,
		"prune", false, "delete all keys not found in github",
	)

	flags.StringVar(
		&arguments.SystemUser,
		"system-user", defaultUser(), "specify the unix system user",
	)

	home := filepath.Dir(os.Getenv("HOME"))
	keys := filepath.Join(home, arguments.SystemUser, ".ssh", "authorized_keys")
	flags.StringVar(
		&arguments.AuthorizedKeys,
		"authorized-keys", "",
		fmt.Sprintf("override the output authorized_keys file (%s)", keys),
	)

	flags.StringVar(
		&arguments.GitHubUser,
		"github-user", "", "specify the github user",
	)

	flags.StringVar(
		&arguments.GitHubAPI,
		"github-api", "https://api.github.com", "specify the GitHub API endpoint",
	)

	_ = flags.Parse(args)
	arguments.usage = flags.Usage
	return arguments
}
