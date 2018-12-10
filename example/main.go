package main

import (
	"github.com/antonmashko/log"
)

func main() {
	log.DefaultLogger = log.New("tag_package_example", log.DebugLevel)
	log.Debugln("hello world")
	log.Copy("tag_package_example").Error("error message")
}
