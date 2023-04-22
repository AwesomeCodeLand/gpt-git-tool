package command

import (
	"fmt"
	"ggt/helper"
	"ggt/tools"

	"github.com/urfave/cli"
)

func Diff() cli.Command {
	return cli.Command{
		Name:  "diff",
		Usage: "get all the differences",
		Action: func(*cli.Context) error {
			_, err := helper.GetConfig()
			if err != nil {
				tools.ErrorDescAndLogin("Diff", err)
			}

			content, err := helper.GetChangeFiles()
			if err != nil {
				tools.ErrorDescAndLogin("Diff", err)
			}

			for name, value := range content {
				fmt.Println("File name: " + name)
				fmt.Println("File content: " + value)
			}
			return nil
		},
	}
}
