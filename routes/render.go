package routes

import (
	"html/template"
	"io"
)

type Template struct {
	tmpls *template.Template
}

func (t *Template) Render(wrt io.Writer, name string, data interface{}) error {
	return t.tmpls.ExecuteTemplate(wrt, name, data)
}
