package command

import (
	"fmt"
	"ggt/helper"
	"ggt/types"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func Diff() cli.Command {
	return cli.Command{
		Name:  "diff",
		Usage: "get all the differences",
		Action: func(c *cli.Context) error {
			var err error
			defer func() {
				if err != nil {
					switch err.(type) {
					case *types.GptError:
						fmt.Fprintf(os.Stderr, "GPT error: %s\n", err.Error())
						os.Exit(1)
					case *types.LoginError:
						fmt.Fprintf(os.Stderr, "Login error: %s\n", err.Error())
						os.Exit(1)
					default:
						fmt.Fprintf(os.Stderr, "Unknown error: %s\n", err.Error())
						os.Exit(1)
					}
				}
			}()

			cfg, err := helper.GetConfig()
			if err != nil {
				// tools.ErrorDescAndLogin("Diff", err)
				return err
			}

			content, err := helper.GetChangeFiles()
			if err != nil {
				// tools.ErrorDescAndLogin("Diff", err)
				return err
			}

			logrus.Debug("---------------------------------")
			ai := helper.NewOpenAI(cfg.Open.Token)
			answer, err := ai.Diff(content)
			if err != nil {
				return err
			}

			logrus.Debug("---------------------------------")
			logrus.Debug("")
			logrus.Info(answer)
			return nil
		},
	}
}
