package http

import (
	"net/http"

	"github.com/ynori7/lilypad/errors"
)

// Response is a basic http response
type Response struct {
	Status      int
	Body        []byte
	RedirectURL string
	Headers     Headers
}

// Headers is a map of HTTP response headers
type Headers map[string]string

// NewResponse returns a new response
func NewResponse(status int, body []byte, redirectUrl string, headers map[string]string) Response {
	return Response{
		Status:      status,
		Body:        body,
		RedirectURL: redirectUrl,
		Headers:     headers,
	}
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
		RedirectURL: path,
	}
}

// WithHeaders returns a new response with the specified HTTP headers added
func (r Response) WithHeaders(h Headers) Response {
	return NewResponse(r.Status, r.Body, r.RedirectURL, mergeMaps(r.Headers, h))
}

// WithMaxAge returns a new response with cache-control headers applied
func (r Response) WithMaxAge(s int64) Response {
	return NewResponse(r.Status, r.Body, r.RedirectURL, mergeMaps(r.Headers, getCacheControlHeaders(s)))
}

func mergeMaps(a map[string]string, b map[string]string) map[string]string {
	c := make(map[string]string)

	for k, v := range a {
		c[k] = v
	}

	for k, v := range b {
		c[k] = v
	}

	return c
}
