package errors

import "sync"

// ErrorType indicates the format that errors should be presented in. They can be Plaintext (default), Markup, or JSON
type ErrorType uint8

const (
	// ErrorTypePlain indicates plaintext errors
	ErrorTypePlain ErrorType = iota
	// ErrorTypeMarkup indicates markup errors (HTML)
	ErrorTypeMarkup
	// ErrorTypeJSON indicates JSON errors
	ErrorTypeJSON
)

type errorConfig struct {
	errorType ErrorType
	template  string
	mutex     sync.RWMutex
}

var defaultErrorConfig errorConfig

// UseMarkupErrors configures the framework to use the provided template for presenting the error.
// For example, an HTML template:
// <html>
// <head></head>
// <body>
// <h1>{{ .Status }}</h1>
// <p>{{ .Message }}</p>
// </body>
// </html>
func UseMarkupErrors(template string) {
	defaultErrorConfig.mutex.Lock()
	defaultErrorConfig.errorType = ErrorTypeMarkup
	defaultErrorConfig.template = template
	defaultErrorConfig.mutex.Unlock()
}

// UseJSONErrors configures the framework to use json for presenting errors. The JSON will consist of a status and a message
func UseJSONErrors() {
	defaultErrorConfig.mutex.Lock()
	defaultErrorConfig.errorType = ErrorTypeJSON
	defaultErrorConfig.mutex.Unlock()
}

// UsePlaintextErrors configures the framework to output the errors as simple plaintext. This is the default
func UsePlaintextErrors() {
	defaultErrorConfig.mutex.Lock()
	defaultErrorConfig.errorType = ErrorTypePlain
	defaultErrorConfig.mutex.Unlock()
}

func getErrorConfig() (ErrorType, string) {
	defaultErrorConfig.mutex.RLock()
	defer defaultErrorConfig.mutex.RUnlock()
	return defaultErrorConfig.errorType, defaultErrorConfig.template
}
