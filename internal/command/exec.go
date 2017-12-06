// Author hoenig

package command

import (
	"github.com/shoenig/ssh-key-sync/internal/config"
	"github.com/shoenig/ssh-key-sync/internal/github"
)

type Execer interface {
	Exec() error
}

func NewExecer(
	loader config.Loader,
	client github.Client,
) Execer {
	return &execer{
		loader: loader,
		client: client,
	}
}

type execer struct {
	loader config.Loader
	client github.Client
}

func (e *execer) Exec() error {
	//opts, err := loader.Load()
	//if err != nil {
	//	fmt.Println("failed to load configuration:", err)
	//	os.Exit(1)
	//}
	//
	//fmt.Println("opts:", opts)

	// client := github.NewClient(nil)
	//keys, err := client.GetKeys(username)
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	//
	//for _, key := range keys {
	//	fmt.Println(key)
	//}
	return nil
}
