package log

import (
	"fmt"
	"io"
	"os"
)

type Logger struct {
	threshold Level
	output io.Writer
}

func New(threshold Level, opts ...Option) *Logger {
	lgr := &Logger{threshold: threshold, output: os.Stdout}

	for _, configFunc := range opts {
		configFunc(lgr)
	}

	return lgr
}

func (l *Logger) Debugf(format string, args ...any) {
	if l.threshold > LevelDebug {
		return
	}

	l.logf(format, args...)
}

func (l *Logger) logf(format string, args ...any) {
	_, _ = fmt.Fprintf(l.output, format+"\n", args...)
}

func (l *Logger) Infof(format string, args ...any) {
	if l.threshold > LeveInfo {
		return
	}

	l.logf(format, args...)
}