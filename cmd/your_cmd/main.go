package main

import (
	"log"
	"os"

	"github.com/your_cmd/cmd"
)

func main() {
	if err := cmd.App().Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
