package logging

type LogLevel int64

const (
	Debug LogLevel = iota
	Info
	Warn
	Error
	Critical
)

type Logger struct {
	LogLevel LogLevel
}

type loggerOptions func(*Logger)

func NewLogger(logLevel LogLevel, options ...loggerOptions) *Logger {
	logger := &Logger{
		LogLevel: logLevel,
	}

	for _, opt := range options {
		opt(logger)
	}

	return logger
}
