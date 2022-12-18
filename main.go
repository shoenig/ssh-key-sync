package main

import (
	"fmt"
	"os"

	"github.com/shoenig/ssh-key-sync/internal/command"
)

func main() {
	if err := command.Start(os.Args); err != nil {
		fmt.Println("[fatal]", err)
		os.Exit(1)
	}
 }
