package log

import (
	"github.com/sirupsen/logrus"
)

// Entry is the final or intermediate logging entry containing all the log fields
type Entry struct {
	logger *logrus.Entry
}

// Debug logs a message at the debug level
func (e *Entry) Debug(args ...interface{}) {
	e.logger.Debug(args...)
}

// Debugf logs a message at the debug level
func (e *Entry) Debugf(format string, args ...interface{}) {
	e.logger.Debugf(format, args...)
}

// Info logs a message at the info level
func (e *Entry) Info(args ...interface{}) {
	e.logger.Info(args...)
}

// Infof logs a message at the info level
func (e *Entry) Infof(format string, args ...interface{}) {
	e.logger.Infof(format, args...)
}

// Warn logs a message at the warn level
func (e *Entry) Warn(args ...interface{}) {
	e.logger.Warn(args...)
}

// Warnf logs a message at the warn level
func (e *Entry) Warnf(format string, args ...interface{}) {
	e.logger.Warnf(format, args...)
}

// Error logs a message at the error level
func (e *Entry) Error(args ...interface{}) {
	e.logger.Error(args...)
}

// Errorf logs a message at the error level
func (e *Entry) Errorf(format string, args ...interface{}) {
	e.logger.Errorf(format, args...)
}

// Fatal logs a message at the fatal level and terminates the program
func (e *Entry) Fatal(args ...interface{}) {
	e.logger.Fatal(args...)
}

// Fatalf logs a message at the fatal level and terminates the program
func (e *Entry) Fatalf(format string, args ...interface{}) {
	e.logger.Fatalf(format, args...)
}

// WithFields adds fields to the logger
func (e *Entry) WithFields(fields Fields) Logger {
	return &Entry{
		logger: e.logger.WithFields(logrus.Fields(fields)),
	}
}
