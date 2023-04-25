package tools

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func Debug(c *cli.Context) {
	logrus.SetLevel(logrus.InfoLevel)
	if c.Bool("debug") {
		logrus.SetLevel(logrus.DebugLevel)
	}
}
