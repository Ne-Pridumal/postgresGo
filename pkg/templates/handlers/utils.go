package handlers

import (
	"html/template"
	"io"
	"path/filepath"
)

type templates struct {
	templates *template.Template
	scripts   *template.JS
}

var path, _ = filepath.Abs(".")

var pagesTemplates = &templates{
	templates: template.Must(template.ParseGlob(path + "/pkg/templates/pages/*/*.html")),
}

func (t *templates) Render(w io.Writer, name string, data any) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
