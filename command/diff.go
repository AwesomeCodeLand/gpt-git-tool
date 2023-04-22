package command

import (
	"ggt/helper"
	"ggt/tools"

	"github.com/sirupsen/logrus"
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

			files, err := helper.GetChangeFiles()
			if err != nil {
				tools.ErrorDescAndLogin("Diff", err)
			}

			logrus.Infof("files: %v", files)

			helper.GetChangeContentWithFile("cli/command.go")
			return nil
		},
	}
}
