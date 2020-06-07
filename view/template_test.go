package view_test

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/ynori7/lilypad/view"
)

type News struct {
	Title       string
	Description string
	Link        string
}

func Test_ExecuteHtmlTemplate(t *testing.T) {
	// given
	news := make([]News, 1)
	news[0] = News{
		Title:       "Some news happened",
		Description: "Lots of really bad stuff happened. Oh no!",
		Link:        "http://www.blah.com/blop",
	}

	// when
	_, err := view.RenderTemplate(`<html>
<head>
</head>
<body>
<section class="news">
	{{ range $i, $val := .News }}
	<div class="news-item">
		<h3><a href="{{ $val.Link }}" target="_blank">{{ $val.Title }}</a></h3>
		<div class="description"><p>{{ $val.Description }}</p></div>
	</div>
	<hr />
	{{ end }}
</section>
</body>
</html>
`, struct {
		News []News
	}{News: news})

	// then
	require.NoError(t, err)
}

func Test_RenderView(t *testing.T) {
	// given
	view.SetLayoutDirectory("../examples/website/view/layout")

	// when
	out, err := view.New("layout", "../examples/website/view/error.gohtml").Render(
		struct {
			Status  int
			Message string
		}{Status: 400, Message: "oops"})

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, out)
}

func Test_RegisterGlobalTemplateFuncs(t *testing.T) {
	// given
	news := make([]News, 1)
	news[0] = News{
		Title: "Some news happened",
	}
	truncate := func(s string) string {
		return s[0:10]
	}

	// when
	view.RegisterGlobalTemplateFuncs(template.FuncMap{
		"truncate": truncate,
	})
	actual, err := view.RenderTemplate(`<html>
<head>
</head>
<body>
<section class="news">
{{ range $i, $val := .News }}<h3>{{ truncate $val.Title }}</h3>{{ end }}
</section>
</body>
</html>
`, struct {
		News []News
	}{News: news})

	// then
	require.NoError(t, err)
	assert.Equal(t, []byte(`<html>
<head>
</head>
<body>
<section class="news">
<h3>Some news </h3>
</section>
</body>
</html>
`), actual)
}
