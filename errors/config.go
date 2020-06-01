package errors

import "sync"

type ErrorType uint8

const (
	ErrorType_Plain ErrorType = iota
	ErrorType_Markup
	ErrorType_Json
)

type errorConfig struct {
	errorType ErrorType
	template string
	mutex sync.RWMutex
}

var defaultErrorConfig errorConfig

/*
 * UseMarkupErrors configures the framework to use the provided template for presenting the error.
 * For example, an HTML template:
 * <html>
 * <head></head>
 * <body>
 * <h1>{{ .Status }}</h1>
 * <p>{{ .Message }}</p>
 * </body>
 * </html>
 */
func UseMarkupErrors(template string) {
	defaultErrorConfig.mutex.Lock()
	defaultErrorConfig.errorType = ErrorType_Markup
	defaultErrorConfig.template = template
	defaultErrorConfig.mutex.Unlock()
}

// UseJsonErrors configures the framework to use json for presenting errors. The JSON will consist of a status and a message
func UseJsonErrors() {
	defaultErrorConfig.mutex.Lock()
	defaultErrorConfig.errorType = ErrorType_Json
	defaultErrorConfig.mutex.Unlock()
}

// UsePlaintextErrors configures the framework to output the errors as simple plaintext. This is the default
func UsePlaintextErrors() {
	defaultErrorConfig.mutex.Lock()
	defaultErrorConfig.errorType = ErrorType_Plain
	defaultErrorConfig.mutex.Unlock()
}

func getErrorConfig() (ErrorType, string) {
	defaultErrorConfig.mutex.RLock()
	defer defaultErrorConfig.mutex.RUnlock()
	return defaultErrorConfig.errorType, defaultErrorConfig.template
}