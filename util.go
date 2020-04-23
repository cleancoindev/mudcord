// util contains various utility "quality of life" functions, that are not specific to anything

package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

// CheckFatal checks if there is a fatal error, and exits accordingly
func CheckFatal(err error) {
	if err != nil {
		logrus.Fatal(err)
		os.Exit(1)
	}
}
