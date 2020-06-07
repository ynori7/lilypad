package errors

import (
	"sync"

	"github.com/ynori7/lilypad/view"
)

// ErrorType indicates the format that errors should be presented in. They can be Plaintext (default), Markup, or JSON
type ErrorType uint8

const (
	// ErrorTypePlain indicates plaintext errors
	ErrorTypePlain ErrorType = iota
	// ErrorTypeMarkup indicates markup errors (HTML)
	ErrorTypeMarkup
	// ErrorTypeMarkupWithLayout indicates markup errors (HTML) with a base layout template
	ErrorTypeMarkupWithLayout
	// ErrorTypeJSON indicates JSON errors
	ErrorTypeJSON
)

type errorConfigWrapper struct {
	conf  errorConfig
	mutex sync.RWMutex
}

type errorConfig struct {
	errorType          ErrorType
	templateRaw        string
	templateWithLayout *view.View
}

var defaultErrorConfig errorConfigWrapper

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
	defaultErrorConfig.conf.errorType = ErrorTypeMarkup
	defaultErrorConfig.conf.templateRaw = template
	defaultErrorConfig.mutex.Unlock()
}

// UseMarkupErrorsWithLayout configures the framework to use templates from files with the specified base layout for presenting the error
func UseMarkupErrorsWithLayout(layout, templateFile string) {
	defaultErrorConfig.mutex.Lock()
	defaultErrorConfig.conf.errorType = ErrorTypeMarkupWithLayout
	defaultErrorConfig.conf.templateWithLayout = view.New(layout, templateFile)
	defaultErrorConfig.mutex.Unlock()
}

// UseJSONErrors configures the framework to use json for presenting errors. The JSON will consist of a status and a message
func UseJSONErrors() {
	defaultErrorConfig.mutex.Lock()
	defaultErrorConfig.conf.errorType = ErrorTypeJSON
	defaultErrorConfig.mutex.Unlock()
}

// UsePlaintextErrors configures the framework to output the errors as simple plaintext. This is the default
func UsePlaintextErrors() {
	defaultErrorConfig.mutex.Lock()
	defaultErrorConfig.conf.errorType = ErrorTypePlain
	defaultErrorConfig.mutex.Unlock()
}

func getErrorConfig() errorConfig {
	defaultErrorConfig.mutex.RLock()
	defer defaultErrorConfig.mutex.RUnlock()
	return defaultErrorConfig.conf
}
