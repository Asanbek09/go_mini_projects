package log_test

import (
	"logger/log"
	"testing"
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

func TestLogger_DebugfInfoErrorf(t *testing.T) {
	type testCase struct {
		level log.Level
		expected string
	}

	tt := map[string]testCase{
		"debug": {
			level: log.LevelDebug,
			expected: debugMessage + "\n" + infoMessage + "\n" + errorMessage + "\n",
		},
		"info": {
			level: log.LeveInfo,
			expected: infoMessage + "\n" + errorMessage + "\n",
		},
		"error": {
			level: log.LevelError,
			expected: errorMessage + "\n",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			tw := &testWriter{}

			testedLogger := log.New(tc.level, log.WithOutput(tw))

			testedLogger.Debugf(debugMessage)
			testedLogger.Infof(infoMessage)
			testedLogger.Errorf(errorMessage)

			if tw.contents != tc.expected {
				t.Errorf("Invalid contents, expected %q, got %q", tc.expected, tw.contents)
			}
		})
	}
}