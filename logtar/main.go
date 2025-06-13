package logtar

import (
	"encoding/json"
	"fmt"
	"os"
)

type LogLevel int64

const (
	Debug LogLevel = iota
	Info
	Warn
	Error
	Critical
)

type Logger struct {
	logLevel LogLevel
}

type loggerOptions func(*Logger)

func WithLogLevel(logLevel LogLevel) loggerOptions {
	return func(l *Logger) {
		l.logLevel = logLevel
	}
}

func NewLogger(options ...loggerOptions) *Logger {
	logger := &Logger{
		logLevel: LogLevel(0),
	}

	for _, opt := range options {
		opt(logger)
	}

	return logger
}

type LogConfig struct {
	Level LogLevel `json:"logLevel"`
	RollingConfig RollingConfig `json:"rollingConfig"`
	FilePrefix string `json:"filePrefix"`
}

func NewLogConfig(path string) (*LogConfig, error) {
	logConfig := LogConfig{}
	file, err := os.Open(path)
	if err != nil {
		return &logConfig, fmt.Errorf("failed to open config file %w", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&logConfig)
	if err != nil {
		return &logConfig, fmt.Errorf("failed to parse config file %w", err)
	}

	return &logConfig, nil
}