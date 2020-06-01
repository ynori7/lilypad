package errors

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ynori7/lilypad/view"
)

// HttpError represents a basic HTTP error with a status and a message
type HttpError struct {
	Status    int    `json:"status"`
	Code      string `json:"code"` //not to be confused with the Status code. This is a string code to identify the error case
	Title     string `json:"title"`
	Message   string `json:"message"`
	Retriable bool   `json:"retriable"`
}

// WithCode sets the error code
func (e HttpError) WithCode(code string) {
	e.Code = code
}

// WithTitle sets the error title
func (e HttpError) WithTitle(title string) {
	e.Title = title
}

// WithRetriable sets the retriable flag on the error
func (e HttpError) WithRetriable(retriable bool) {
	e.Retriable = retriable
}

// New returns a new http error
func New(status int, code string, title string, message string, retriable bool) HttpError {
	return HttpError{
		Status:    status,
		Code:      code,
		Message:   message,
		Title:     title,
		Retriable: retriable,
	}
}

// InternalServerError returns a 500 error
func InternalServerError(message string) HttpError {
	return HttpError{
		Status:    http.StatusInternalServerError,
		Message:   message,
		Retriable: true,
	}
}

// BadRequestError returns a 400 error
func BadRequestError(message string) HttpError {
	return HttpError{
		Status:    http.StatusBadRequest,
		Message:   message,
		Retriable: false,
	}
}

// NotFoundError returns a 404 error
func NotFoundError(message string) HttpError {
	return HttpError{
		Status:    http.StatusNotFound,
		Message:   message,
		Retriable: false,
	}
}

// Write returns a string representation of the error based on the global configuration. In case of failure, an error is returned
func (e HttpError) Write() (string, error) {
	errorType, template := getErrorConfig()

	switch errorType {
	case ErrorType_Markup:
		out, err := view.RenderTemplate(template, e)
		if err != nil {
			return "", err
		}
		return out, nil

	case ErrorType_Json:
		out, err := json.Marshal(e)
		if err != nil {
			return "", err
		}
		return string(out), nil

	default:
		return fmt.Sprintf("%d %s", e.Status, e.Message), nil
	}
}
