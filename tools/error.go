package tools

import (
	"os"

	"github.com/sirupsen/logrus"
)

func ErrorDescAndLogin(preFix string, err error) {
	logrus.Errorf("%s: %v \n", preFix, err)
	logrus.Errorln("Please Login First")
	os.Exit(1)
}
