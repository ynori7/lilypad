package log

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

const ipHeader = "X-FORWARDED_FOR"

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	SetLevel(LevelDebug)
}

// UseJSONFormatter sets the log format to JSON
func UseJSONFormatter() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

// UseTextFormatter sets the log format to plain text
func UseTextFormatter() {
	logrus.SetFormatter(&logrus.TextFormatter{})
}

// Fields is a map of attributes to be included when logging
type Fields map[string]interface{}

// Logger is an entity which handles logging
type Logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})

	Error(args ...interface{})
	Errorf(format string, args ...interface{})

	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})

	Info(args ...interface{})
	Infof(format string, args ...interface{})

	Warn(args ...interface{})
	Warnf(format string, args ...interface{})

	WithFields(fields Fields) Logger
}

// Debug logs a message at the debug level
func Debug(args ...interface{}) {
	logrus.Debug(args...)
}

// Debugf logs a message at the debug level
func Debugf(format string, args ...interface{}) {
	logrus.Debugf(format, args...)
}

// Info logs a message at the info level
func Info(args ...interface{}) {
	logrus.Info(args...)
}

// Infof logs a message at the info level
func Infof(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

// Warn logs a message at the warn level
func Warn(args ...interface{}) {
	logrus.Warn(args...)
}

// Warnf logs a message at the warn level
func Warnf(format string, args ...interface{}) {
	logrus.Warnf(format, args...)
}

// Error logs a message at the error level
func Error(args ...interface{}) {
	logrus.Error(args...)
}

// Errorf logs a message at the error level
func Errorf(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}

// Fatal logs a message at the fatal level
func Fatal(args ...interface{}) {
	logrus.Fatal(args...)
}

// Fatalf logs a message at the fatal level
func Fatalf(format string, args ...interface{}) {
	logrus.Fatalf(format, args...)
}

// WithFields creates a logger with log fields added to it
func WithFields(fields Fields) Logger {
	return &Entry{
		logger: logrus.WithFields(logrus.Fields(fields)),
	}
}

// WithRequest adds fields from the http request to the logger
func WithRequest(r *http.Request) Logger {
	return &Entry{
		logger: logrus.WithFields(logrus.Fields{
			"ClientIp": getIPFromRequest(r),
		}),
	}
}

func getIPFromRequest(r *http.Request) string {
	forwardedIP := r.Header.Get(ipHeader)
	if forwardedIP != "" {
		return forwardedIP
	}

	return r.RemoteAddr
}
