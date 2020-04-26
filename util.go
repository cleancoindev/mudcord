// util contains various utility "quality of life" functions, that are not specific to anything

package main

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"time"
)

// CheckFatal checks if there is a fatal error, and exits accordingly
func CheckFatal(err error) {
	if err != nil {
		logrus.Fatal(err)
		os.Exit(1)
	}
}

func Serializer() {
	for {
		b, err := json.Marshal(Data)
		CheckFatal(err)
		ioutil.WriteFile("data.json", b, 0644)
		time.Sleep(10 * time.Second)
	}
}
