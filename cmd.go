package discord

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

var conf Conf

func init() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	// initialize the configuration
	conf = Conf{
		Id:   "run",
		Dir:  filepath.Join(homedir, ".config", "discord-go"),
		File: "config.json",
	}

	if err := conf.Init(); err != nil {
		log.Fatal(err)
	}
}

func App() *cli.App {
	app := &cli.App{
		Name:                 "Run any command on any repo",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name: "hello",
				Action: func(c *cli.Context) error {
					fmt.Println("Hello world!")
					return nil
				},
			},
		},
	}
	return app
}
