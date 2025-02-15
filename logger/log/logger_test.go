package log_test

import (
	"logger/log"
)
func ExampleLogger_Debugf() {
	debugLogger := log.New(log.LevelDebug)
	debugLogger.Debugf("Hello, %s", "world")
}