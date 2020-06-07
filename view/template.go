package view

import (
	"bufio"
	"bytes"
	"html/template"
	"path/filepath"
)

// RenderTemplate generates output with the given raw string template and parameters. If the execution fails, an error is returned.
func RenderTemplate(templateToRender string, data interface{}) ([]byte, error) {
	t := template.Must(template.New("html").
		Funcs(getFuncMap()).
		Parse(templateToRender))

	var b bytes.Buffer
	w := bufio.NewWriter(&b)

	err := t.Execute(w, data)
	if err != nil {
		return nil, err
	}

	w.Flush()
	return b.Bytes(), nil
}

// View is the configuration for a particular view with a template and its base layout
type View struct {
	template *template.Template
	layout   string
}

// New creates a new view with the given layout name (created with a "define" block in the layout directory) and the specified template files
func New(layout string, files ...string) *View {
	files = append(layoutFiles(), files...)

	t := template.Must(template.New("html").
		Funcs(getFuncMap()).
		ParseFiles(files...))

	return &View{
		template: t,
		layout:   layout,
	}
}

// Render generates HTML output from the view
func (v *View) Render(data interface{}) ([]byte, error) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)

	err := v.template.ExecuteTemplate(w, v.layout, data)
	if err != nil {
		return nil, err
	}

	w.Flush()
	return b.Bytes(), nil
}

func layoutFiles() []string {
	files, err := filepath.Glob(getLayoutDir() + "/*.gohtml")
	if err != nil {
		panic(err)
	}
	return files
}
