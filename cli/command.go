package main

// use urfave/cli generate command
import (
	"fmt"

	"github.com/urfave/cli"
)

// command parses the command line arguments and returns the command
func command() *cli.App {
	return &cli.App{
		Commands: []cli.Command{
			login(),
			diff(),
		},
	}
}

// diff a command line tool to get all the differences
func diff() cli.Command {
	return cli.Command{
		Name:  "diff",
		Usage: "get all the differences",
		Action: func(*cli.Context) error {
			fmt.Println("Hello friend!")
			return nil
		},
	}
}

func login() cli.Command {
	return cli.Command{
		Name:  "login",
		Usage: "login to openai, type your openai secret key. you can get it from https://beta.openai.com/account/api-keys. it will be saved to ~/.ggt/config.json",
		Action: func(*cli.Context) error {
			fmt.Println("Hello friend!")
			return nil
		},
	}
}
