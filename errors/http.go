package errors

import (
	"encoding/json"
	"fmt"

	"github.com/ynori7/lilypad/view"
)

// HttpError represents a basic HTTP error with a status and a message
type HttpError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
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
