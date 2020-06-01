package log

import (
	"github.com/sirupsen/logrus"
)

type Entry struct {
	logger *logrus.Entry
}

func (e *Entry) Debug(args ...interface{}) {
	e.logger.Debug(args...)
}

func (e *Entry) Debugf(format string, args ...interface{}) {
	e.logger.Debugf(format, args...)
}

func (e *Entry) Info(args ...interface{}) {
	e.logger.Info(args...)
}

func (e *Entry) Infof(format string, args ...interface{}) {
	e.logger.Infof(format, args...)
}

func (e *Entry) Warn(args ...interface{}) {
	e.logger.Warn(args...)
}

func (e *Entry) Warnf(format string, args ...interface{}) {
	e.logger.Warnf(format, args...)
}

func (e *Entry) Error(args ...interface{}) {
	e.logger.Error(args...)
}

func (e *Entry) Errorf(format string, args ...interface{}) {
	e.logger.Errorf(format, args...)
}

func (e *Entry) Fatal(args ...interface{}) {
	e.logger.Fatal(args...)
}

func (e *Entry) Fatalf(format string, args ...interface{}) {
	e.logger.Fatalf(format, args...)
}

func (e *Entry) WithFields(fields Fields) Logger {
	return &Entry{
		logger: e.logger.WithFields(logrus.Fields(fields)),
	}
}
