package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	app := command()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
