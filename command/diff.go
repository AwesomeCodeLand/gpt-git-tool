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

			fmt.Println("---------------------------------")
			for k, v := range content {
				fmt.Println(k)
				fmt.Println(v)
				fmt.Println("---------------------------------")
			}
			// ai := helper.NewOpenAI(cfg.Open.Token)
			// answer, err := ai.Diff(content)
			// if err != nil {
			// 	tools.ErrorDescAndLogin("Diff", err)
			// 	os.Exit(1)
			// }

			// fmt.Println("---------------------------------")
			// fmt.Println("")
			// fmt.Println(answer)
			return nil
		},
	}
}
