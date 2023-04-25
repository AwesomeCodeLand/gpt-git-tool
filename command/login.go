package command

import (
	"ggt/config"
	"ggt/helper"
	"ggt/tools"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func Login() cli.Command {
	return cli.Command{
		Name:  "login",
		Usage: "login to openai, type your openai secret key. you can get it from https://beta.openai.com/account/api-keys. it will be saved to ~/.ggt/config.json",
		Action: func(cx *cli.Context) error {
			tools.Debug(cx)

			key := cx.Args().Get(0)

			err := helper.SaveConfig(config.Cfg{
				Open: config.OpenAI{
					Token: key,
				},
			})
			if err != nil {
				logrus.Errorf("login error: %v", err)
				os.Exit(1)
			}

			logrus.Info("Login success")
			return nil
		},
	}
}
