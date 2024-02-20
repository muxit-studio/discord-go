package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// var (
// 	conf = Conf{
// 		Id:   "run",
// 		Dir:  "/home/odas0r/github.com/odas0r/configs",
// 		File: "config.json",
// 	}
// )

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
