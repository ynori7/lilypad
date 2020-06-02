package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_WriteHtmlError(t *testing.T) {
	// given
	template := `<html>
<head></head>
<body>
<h1>{{ .Status }}</h1>
<p>{{ .Message }}</p>
</body>
</html>
`
	err := HttpError{
		Status:  500,
		Message: "Something went wrong",
	}

	// when
	UseMarkupErrors(template)
	actual, e := err.Write()

	// then
	require.NoError(t, e)
	assert.Equal(t, `<html>
<head></head>
<body>
<h1>500</h1>
<p>Something went wrong</p>
</body>
</html>
`, string(actual))
}

func Test_WritePlaintextError(t *testing.T) {
	//  given
	err := HttpError{
		Status:  500,
		Message: "Something went wrong",
	}

	// when
	UsePlaintextErrors()
	actual, e := err.Write()

	// then
	require.NoError(t, e)
	assert.Equal(t, "500 Something went wrong", string(actual))
}

func Test_WriteJsonError(t *testing.T) {
	//  given
	err := HttpError{
		Status:    500,
		Code:      "BROKEN",
		Title:     "Uh oh",
		Message:   "Something went wrong",
		Retriable: true,
	}

	// when
	UseJSONErrors()
	actual, e := err.Write()

	// then
	require.NoError(t, e)
	assert.JSONEq(t, `{"status":500,"code":"BROKEN","title":"Uh oh","message":"Something went wrong","retriable":true}`, string(actual))
}

func Test_WriteJsonError_WithExtraFields(t *testing.T) {
	//  given
	err := NotFoundError("page not found").WithTitle("missing").WithCode("NOT_FOUND")

	// when
	UseJSONErrors()
	actual, e := err.Write()

	// then
	require.NoError(t, e)
	assert.JSONEq(t, `{"status":404,"code":"NOT_FOUND","title":"missing","message":"page not found","retriable":false}`, string(actual))
}
