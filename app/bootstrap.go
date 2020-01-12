package app

import (
	"fmt"
	"os"

	"github.com/jinzhu/configor"
	"github.com/sirupsen/logrus"
)

func init() {
	configor.Load(&Config, "conf/config.yml")
	fmt.Printf("config: %#v", Config)
	registLog()
}

func registLog() {
	var file *os.File
	var err error
	// Log as JSON instead of the default ASCII formatter.
	file, err = os.OpenFile("logs/golang.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	if err != nil {
		logrus.SetOutput(os.Stdout)
	} else {
		logrus.SetOutput(file)
	}
}
