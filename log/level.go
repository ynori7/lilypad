package log

import "github.com/sirupsen/logrus"

// Level is an enum representing the different log levels
type Level int

const (
	LevelFatal Level = iota
	LevelError
	LevelWarn
	LevelInfo
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
