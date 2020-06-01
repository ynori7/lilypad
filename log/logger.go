package log

import (
	"net/http"
	
	"github.com/sirupsen/logrus"
)

const IpHeader = "X-FORWARDED_FOR"

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	SetLevel(LevelDebug)
}

type Fields map[string]interface{}

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

func Debug(args ...interface{}) {
	logrus.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	logrus.Debugf(format, args...)
}

func Info(args ...interface{}) {
	logrus.Info(args...)
}

func Infof(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

func Warn(args ...interface{}) {
	logrus.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	logrus.Warnf(format, args...)
}

func Error(args ...interface{}) {
	logrus.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}

func Fatal(args ...interface{}) {
	logrus.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	logrus.Fatalf(format, args...)
}

func WithFields(fields Fields) Logger {
	return &Entry{
		logger: logrus.WithFields(logrus.Fields(fields)),
	}
}

// WithRequest adds fields from the http request to the logger
func WithRequest(r *http.Request) Logger {
	return &Entry{
		logger: logrus.WithFields(logrus.Fields{
			"ClientIp": getIpFromRequest(r),
		}),
	}
}

func getIpFromRequest(r *http.Request) string {
	forwardedIp := r.Header.Get(IpHeader)
	if forwardedIp != "" {
		return forwardedIp
	}

	return r.RemoteAddr
}
