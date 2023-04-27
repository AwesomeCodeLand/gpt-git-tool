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
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "file",
				Value: "",
				Usage: "Specify the file to summary",
			},
		},
		Action: func(c *cli.Context) error {
			specifyFile := c.String("file")

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

			if specifyFile != "" {
				if _, ok := content[specifyFile]; !ok {
					logrus.Errorf("File %s not found", specifyFile)
					os.Exit(1)
				}
				content = map[string]string{
					specifyFile: content[specifyFile],
				}
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
