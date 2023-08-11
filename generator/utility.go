package generator

import (
	"strings"
	"text/template"
)

func GetFuncs() template.FuncMap {

	return template.FuncMap{
		"capitalize": func(s string) string {
			return strings.Title(s)
		},
	}
}
