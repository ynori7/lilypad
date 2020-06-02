package handler

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
				RedirectUrl: "http://www.blah.com",
			},
		},
		"temporary redirect": {
			actual: RedirectResponse("http://www.blah.com", false),
			expected: Response{
				Status:      302,
				RedirectUrl: "http://www.blah.com",
			},
		},
	}

	for testcase, testresults := range testdata {
		// then
		assert.Equal(t, testresults.expected, testresults.actual, testcase)
	}
}
