package logging

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

type LogLevel int64

const (
	Debug LogLevel = iota
	Info
	Warn
	Error
	Critical
)

// Logger is used for writing information to external files.
type Logger struct {
	LogLevel      LogLevel
	filePrefix    string
	sizeThreshold int

	logFile *os.File
}

func NewLogger(logLevel LogLevel, filePrefix string, sizeThreshold int) *Logger {
	return &Logger{
		LogLevel:      logLevel,
		filePrefix:    filePrefix,
		sizeThreshold: sizeThreshold,
	}
}

func (l *Logger) Open() error {
	err := os.MkdirAll("log", 0755)
	if err != nil {
		return err
	}

	file, err := os.Create("log/" + l.generateLogFileName())
	if err != nil {
		return err
	}

	l.logFile = file

	return nil
}

func (l *Logger) Close() error {
	err := l.logFile.Close()
	if err != nil {
		return err
	}
	return nil
}

func (l *Logger) Debug(msg string) {
	l.log(msg, Debug)
}

func (l *Logger) Info(msg string) {
	l.log(msg, Info)
}

func (l *Logger) Warn(msg string) {
	l.log(msg, Warn)
}

func (l *Logger) Error(msg string) {
	l.log(msg, Error)
}

func (l *Logger) Critical(msg string) {
	l.log(msg, Critical)
}

func (logLevel LogLevel) String() string {
	switch logLevel {
	case Debug:
		return "DEBUG"
	case Info:
		return "INFO"
	case Warn:
		return "WARN"
	case Error:
		return "ERROR"
	case Critical:
		return "CRITICAL"
	default:
		return "UNKNOWN"
	}
}

func (l *Logger) generateLogFileName() string {
	formatedDatetime := time.Now().Format("2006-01-02T15:04:05Z")
	return fmt.Sprintf("%s%s", l.filePrefix, formatedDatetime)
}

func (l *Logger) log(msg string, logLevel LogLevel) error {
	if l.LogLevel > logLevel {
		return fmt.Errorf("logger level (%s) not compatable with argument level (%s)", l.LogLevel, logLevel)
	}

	_, file, line, ok := runtime.Caller(1)
	if !ok {
		return fmt.Errorf("could not retrive caller info")
	}

	err := l.rollingCheck()
	if err != nil {
		return fmt.Errorf("writing to log file: %s", err)
	}

	isoNow := time.Now().UTC().Format(time.RFC3339)

	_, err = l.logFile.WriteString(fmt.Sprintf("[%s] %s %s:%d (%s)\n", isoNow, logLevel, file, line, msg))
	if err != nil {
		return err
	}

	return nil
}

func (l *Logger) rollingCheck() error {
	fileInfo, err := l.logFile.Stat()
	if err != nil {
		return fmt.Errorf("rolling check: %s", err)
	}

	if fileInfo.Size() >= int64(l.sizeThreshold) {
		l.Close()
		l.Open()
	}

	return nil
}
