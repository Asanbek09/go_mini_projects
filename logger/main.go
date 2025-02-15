package main

import (
	"os"

	"logger/log"
)

func main() {
	lgr := log.New(log.LeveInfo, log.WithOutput(os.Stdout))

	lgr.Infof("A little copying is better than a little dependency")
}