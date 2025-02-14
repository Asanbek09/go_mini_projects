package log

import (
	"fmt"
	"io"
)

type Logger struct {
	threshold Level
	output io.Writer
}

func New(threshold Level, output io.Writer) *Logger {
	return &Logger{
		threshold: threshold,
		output: output,
	}
}

func (l *Logger) Debugf(format string, args ...any) {
	if l.threshold > LevelDebug {
		return
	}

	_, _ = fmt.Printf(format+"\n", args...)
}

func (l *Logger) Infof(format string, args ...any) {}