package log_test

import (
	"logger/log"
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