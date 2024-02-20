package main

import (
	"log"
	"os"

	cmd "github.com/muxit-studio/discord-go"
)

func main() {
	if err := cmd.App().Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
