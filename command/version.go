package command

import (
	"fmt"

	"github.com/urfave/cli"
)

var version string = "0.1.0"
var commit string

func Version() cli.Command {
	return cli.Command{
		Name:  "version",
		Usage: "show the version",
		Action: func(c *cli.Context) error {
			fmt.Printf("Version: %s %s\n", version, commit)
			return nil
		},
	}
}
