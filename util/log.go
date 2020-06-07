package util

import (
	"os"
	"log"
)

var Logger *log.Logger

func init() {
	Logger = log.New(os.Stdout, "", log.Ldate | log.Ltime | log.Lshortfile)
}
