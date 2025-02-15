package log_test

import (
	"logger/log"
)

const (
	debugMessage = "Why write I still all one, ever the same"
	infoMessage = "And keep invention in a noted weed"
	errorMessage = "That every word doth almost tell my name"
)

type testWriter struct {
	contents string
}

func ExampleLogger_Debugf() {
	debugLogger := log.New(log.LevelDebug)
	debugLogger.Debugf("Hello, %s", "world")
}

func (tw *testWriter) Write(p []byte) (n int, err error) {
	tw.contents = tw.contents + string(p)
	return len(p), nil
}