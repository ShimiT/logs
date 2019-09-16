package logger

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

const (
	// app log level
	AppLogLevel = "APP_LOG"
	defLevel    = "error"
)

type Logger struct {
	label   string
	version string

	loggerImpl *logrus.Logger
}

// NewLogger - init logger
func NewLogger(level string, version string) *Logger {

	lvl := logLevel(level)
	logger := &logrus.Logger{
		Out:       os.Stdout,
		Level:     lvl,
		//Formatter: &logrus.TextFormatter{},
		Formatter: &prefixed.TextFormatter{
			DisableColors:   false,
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
			ForceFormatting: true,
		},
	}
	return &Logger{
		version:    version,
		loggerImpl: logger,
	}
}

// GetLogLevel - Get level from env
func getLogLevel() string {
	lvl, _ := os.LookupEnv(AppLogLevel)
	if lvl != "" {
		return lvl
	}
	return defLevel
}

func logLevel(lvl string) logrus.Level {

	switch lvl {
	case "debug":
		// Used for tracing
		return logrus.DebugLevel
	case "info":
		return logrus.InfoLevel
	case "error":
		return logrus.ErrorLevel
	case "warn":
		return logrus.WarnLevel
	case "fatal":
		return logrus.FatalLevel
	case "panic":
		return logrus.PanicLevel
	default:
		panic(fmt.Sprintf("the specified %s log level is not supported", lvl))
	}
}

func (logger *Logger) SetLevel(level string) {
	lvl := logLevel(level)
	logger.loggerImpl.SetLevel(lvl)
}

func (logger *Logger) Tracef(format string, args ...interface{}) {
	logger.loggerImpl.Tracef(format, args...)
}

func (logger *Logger) Debugf(format string, args ...interface{}) {
	logger.loggerImpl.Debugf(format, args...)
}
func (logger *Logger) Infof(format string, args ...interface{}) {
	logger.loggerImpl.Infof(format, args...)
}

func (logger *Logger) Warnf(format string, args ...interface{}) {
	logger.loggerImpl.Warnf(format, args...)
}

func (logger *Logger) Errorf(format string, args ...interface{}) {
	logger.loggerImpl.Errorf(format, args...)
}

func (logger *Logger) Fatalf(format string, args ...interface{}) {
	logger.loggerImpl.Fatalf(format, args...)
}

func (logger *Logger) Panicf(format string, args ...interface{}) {
	logger.loggerImpl.Panicf(format, args...)
}

func (logger *Logger) Trace(args ...interface{}) {
	logger.loggerImpl.Trace(args...)
}

func (logger *Logger) Debug(args ...interface{}) {
	logger.loggerImpl.Debug(args...)
}

func (logger *Logger) Info(args ...interface{}) {
	logger.loggerImpl.Info(args...)
}

func (logger *Logger) Warn(args ...interface{}) {
	logger.loggerImpl.Warn(args...)
}

func (logger *Logger) Error(args ...interface{}) {
	logger.loggerImpl.Error(args...)
}

func (logger *Logger) Fatal(args ...interface{}) {
	logger.loggerImpl.Fatal(args...)
}

func (logger *Logger) Panic(args ...interface{}) {
	logger.loggerImpl.Panic(args...)
}
