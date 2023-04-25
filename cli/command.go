package main

// use urfave/cli generate command
import (
	cmd "ggt/command"

	"github.com/urfave/cli"
)

// command parses the command line arguments and returns the command
func command() *cli.App {
	return &cli.App{
		Name:  "ggt",
		Usage: "A git tools base on gpt",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "debug,d",
				Usage: "output debug info",
			},
		},
		Commands: []cli.Command{
			cmd.Login(),
			cmd.Diff(),
		},
	}
}
