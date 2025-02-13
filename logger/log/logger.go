package log

type Logger struct {
	threshold Level
}

func New(threshold Level) *Logger {
	return &Logger{
		threshold: threshold,
	}
}

func (l *Logger) Debugf(format string, args ...any) {}

func (l *Logger) Infof(format string, args ...any) {}