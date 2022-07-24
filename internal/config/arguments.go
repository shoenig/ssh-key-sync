package config

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

type Arguments struct {
	Verbose bool

	SystemUser     string
	SystemHome     string
	AuthorizedKeys string

	GitHubUser string
	GitHubAPI  string
}

func ParseArguments() Arguments {
	var args Arguments

	flag.BoolVar(
		&args.Verbose,
		"verbose", false, "print verbose logging",
	)

	flag.StringVar(
		&args.SystemUser,
		"system-user", os.Getenv("USER"), "specify the unix system user",
	)

	home := filepath.Dir(os.Getenv("HOME"))
	keys := filepath.Join(home, args.SystemUser, ".ssh", "authorized_keys")
	flag.StringVar(
		&args.AuthorizedKeys,
		"authorized-keys", "",
		fmt.Sprintf("override the output authorized_keys file (%s)", keys),
	)

	flag.StringVar(
		&args.GitHubUser,
		"github-user", "", "specify the github user",
	)

	flag.StringVar(
		&args.GitHubAPI,
		"github-api", "https://api.github.com", "specify the GitHub API endpoint",
	)

	flag.Parse()
	return args
}
