package http

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/ynori7/lilypad/errors"
)

func Test_Response(t *testing.T) {
	// given
	errors.UsePlaintextErrors()

	// when
	testdata := map[string]struct {
		actual   Response
		expected Response
	}{
		"success": {
			actual: SuccessResponse([]byte("it works!")),
			expected: Response{
				Status: 200,
				Body:   []byte("it works!"),
			},
		},
		"internal error": {
			actual: ErrorResponse(errors.InternalServerError("something went wrong")),
			expected: Response{
				Status: 500,
				Body:   []byte("500 something went wrong"),
			},
		},
		"not found error": {
			actual: ErrorResponse(errors.NotFoundError("page not found")),
			expected: Response{
				Status: 404,
				Body:   []byte("404 page not found"),
			},
		},
		"bad request error": {
			actual: ErrorResponse(errors.BadRequestError("you messed up")),
			expected: Response{
				Status: 400,
				Body:   []byte("400 you messed up"),
			},
		},
		"permanent redirect": {
			actual: RedirectResponse("http://www.blah.com", true),
			expected: Response{
				Status:      301,
				RedirectURL: "http://www.blah.com",
			},
		},
		"temporary redirect": {
			actual: RedirectResponse("http://www.blah.com", false),
			expected: Response{
				Status:      302,
				RedirectURL: "http://www.blah.com",
			},
		},
	}

	for testcase, testresults := range testdata {
		// then
		assert.Equal(t, testresults.expected, testresults.actual, testcase)
	}
}

func Test_WithCacheHeaders(t *testing.T) {
	// given
	errors.UsePlaintextErrors()

	// when
	resp := SuccessResponse([]byte("it works!")).WithMaxAge(300)

	// then
	assert.Equal(t, 200, resp.Status)
	require.NotEmpty(t, resp.Headers)
	assert.NotEmpty(t, resp.Headers[cacheControlHeader])
	assert.NotEmpty(t, resp.Headers[expiresHeader])
	assert.NotEmpty(t, resp.Headers[lastModifiedHeader])
}

func Test_ErrorResponse_WithWriteFailure(t *testing.T) {
	// given
	errors.UseMarkupErrors("<html><bodyblah") //invalid template

	// when
	r := ErrorResponse(errors.BadRequestError("you messed up"))

	// then
	assert.Equal(t, 500, r.Status, "it should be a 500 even though we returned a 400")
	assert.True(t, strings.HasPrefix(string(r.Body), "html/template:"))
}
