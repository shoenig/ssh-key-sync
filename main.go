// Author hoenig

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/shoenig/ssh-key-sync/internal/github"
)

func main() {
	username := parseArgs()

	client := github.NewClient(nil)

	keys, err := client.GetKeys(username)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, key := range keys {
		fmt.Println(key)
	}
}

func parseArgs() string {
	var username string
	flag.StringVar(&username, "username", "", "the github username")
	flag.Parse()
	return username
}
