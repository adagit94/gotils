package strings

import (
	"text/template"
)

func CreateTemplate(name string, tempText string) *template.Template {
	return template.Must(template.New(name).Parse(tempText))
}