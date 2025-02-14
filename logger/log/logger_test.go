package log_test

import (
	"os"
	"logger/log"
)
func ExampleLogger_Debugf() {
	debugLogger := log.New(log.LevelDebug, os.Stdout)
	debugLogger.Debugf("Hello, %s", "world")
}