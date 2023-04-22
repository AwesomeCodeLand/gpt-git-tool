package tools

import (
	"os"

	"github.com/sirupsen/logrus"
)

func ErrorDescAndLogin(preFix string, err error) {
	logrus.Printf("%s: %v", preFix, err)
	logrus.Println("Please Login First")
	os.Exit(1)
}
