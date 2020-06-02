package log

import "github.com/sirupsen/logrus"

// Level is an enum representing the different log levels
type Level int

const (
	// LevelFatal indicates the fatal log level. Logs at this level will terminate the program,
	// so this should only be used when the whole process is unrecoverable.
	LevelFatal Level = iota
	// LevelError indicates the error log level. Errors represent failures to process a request which could not be
	// handled gracefully.
	LevelError
	// LevelWarn indicates the warn log level. Warnings represent errors which can be handled gracefully.
	LevelWarn
	// LevelInfo indicates the info level. This is used for useful information about the request being handled.
	LevelInfo
	// LevelDebug indicates the debug level. This level is used for fine-grained logging to help with debugging.
	LevelDebug
)

const defaultLogLevel = logrus.DebugLevel

// SetLevel sets the log level globally
func SetLevel(level Level) {
	if logrusLevel, ok := levelMap[level]; ok {
		logrus.SetLevel(logrusLevel)
	} else {
		logrus.SetLevel(defaultLogLevel)
	}
}

var levelMap = map[Level]logrus.Level{
	LevelFatal: logrus.FatalLevel,
	LevelError: logrus.ErrorLevel,
	LevelWarn:  logrus.WarnLevel,
	LevelInfo:  logrus.InfoLevel,
	LevelDebug: logrus.DebugLevel,
}
