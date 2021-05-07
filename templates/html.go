package templates

import (
	"github.com/aaronland/go-http-maps/templates/html"
	"html/template"
)

func LoadHTMLTemplates() (*template.Template, error) {

	t := template.New("map").Funcs(template.FuncMap{
		//
	})

	return t.ParseFS(html.FS, "*.html")
}
