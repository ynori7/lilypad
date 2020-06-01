package view

import (
	"bufio"
	"bytes"
	"html/template"
)

// RenderTemplate generates output with the given template and parameters. If the execution fails, an error is returned.
func RenderTemplate(templateToRender string, data interface{}) (string, error) {
	t := template.Must(template.New("html").
		Funcs(getFuncMap()).
		Parse(templateToRender))

	var b bytes.Buffer
	w := bufio.NewWriter(&b)

	err := t.Execute(w, data)
	if err != nil {
		return "", err
	}

	w.Flush()
	return b.String(), nil
}