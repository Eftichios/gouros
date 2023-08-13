package generator

import (
	"strings"
	"text/template"
	"toncode/gouros/parser"
	"toncode/gouros/transform"
)



func GetFuncs() template.FuncMap {

	return template.FuncMap{
		"capitalize": func(s string) string {
			return strings.Title(s)
		},
		"abbrRepo": func(s string) string {
			return string(s[0]) + "r"
		},
		"abbrServ": func(s string) string {
			return string(s[0]) + "s"
		},
		"abbrContr": func(s string) string {
			return string(s[0]) + "c"
		},
		"fnName": func(r string, m string) string {
			return transform.FnNameTransformer(r, m)
		},
		"createRoute": func(base string, endpoint string) string {
			if endpoint != "/" {
				return base + endpoint
			}
			return base
		},
		"params": func(m *parser.Model, r *parser.Route) string {
			return transform.ParamTransformer(m, r)
		},
	}
}
