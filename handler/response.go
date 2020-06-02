package handler

import (
	"net/http"

	"github.com/ynori7/lilypad/errors"
)

// Response is a basic http response
type Response struct {
	Status      int
	Body        []byte
	RedirectUrl string
}

// SuccessResponse returns a successful http response
func SuccessResponse(body []byte) Response {
	return Response{
		Status: 200,
		Body:   body,
	}
}

// ErrorResponse returns a non-successful response
func ErrorResponse(err errors.HttpError) Response {
	body, e := err.Write()
	if e != nil {
		return Response{
			Status: http.StatusInternalServerError,
			Body:   []byte(e.Error()),
		}
	}

	return Response{
		Status: err.Status,
		Body:   body,
	}
}

// RedirectResponse returns a redirect response
func RedirectResponse(path string, permanent bool) Response {
	status := http.StatusFound
	if permanent {
		status = http.StatusMovedPermanently
	}
	return Response{
		Status:      status,
		RedirectUrl: path,
	}
}
