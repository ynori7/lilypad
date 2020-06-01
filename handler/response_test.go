package handler

import (
	"fmt"
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
			actual: SuccessResponse("it works!"),
			expected: Response{
				Status: 200,
				Body:   "it works!",
			},
		},
		"error": {
			actual: ErrorResponse(500, fmt.Errorf("uh oh")),
			expected: Response{
				Status: 500,
				Body:   "500 uh oh",
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
