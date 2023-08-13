package generator

import (
	"fmt"
	"strings"
	"text/template"
	"toncode/gouros/parser"
)

func fnNameTransformer(route string, method string) string {
	// map routes and methods to function names
	if route == "/" {
		// can potentially use a map here
		switch method {
		case "GET":
			return "GetAll"
		case "POST":
			return "Create"
		case "PUT":
			return "Update"
		case "DELETE":
			return "Delete"

		}
	} else if strings.Contains(route, "/:") {
		params := strings.Split(route, "/:")

		// remove empty string from index 0
		params = params[1:]
		fnName := ""
		isFirst := true
		if len(params) > 1 {
			for _, param := range params {
				if isFirst {
					fnName += strings.Title(param)
					isFirst = false
				} else {
					fnName += "And" + strings.Title(param)
				}
			}
		} else {
			fnName += strings.Title(params[0])
		}

		switch method {
		case "GET":
			return "GetBy" + fnName
		case "PUT":
			return "UpdateBy" + fnName
		case "DELETE":
			return "DeleteBy" + fnName

		}

	}

	return "ReplaceMe"

}

func paramTransformer(model *parser.Model, route *parser.Route) string {
	// create map for param names to types
	attrMap := make(map[string]string)

	for _, attr := range model.Attributes {
		attrMap[attr.Column] = attr.Type // anything else to do here?
	}

	// extract url params
	params := strings.Split(route.Endpoint, "/:")

	// remove empty string from index 0
	params = params[1:]

	result := ""
	for i, param := range params {
		result += fmt.Sprintf("%s %s", param, attrMap[param])
		if i != len(params)-1 {
			result += ", "
		}

	}

	if route.Method == "PUT" || route.Method == "POST" {
        if result != "" {
            result += ", "
        }
		result += fmt.Sprintf("%s *model.%s", model.Table, strings.Title(model.Table))
	}

	return result
}

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
			return fnNameTransformer(r, m)
		},
		"createRoute": func(base string, endpoint string) string {
			if endpoint != "/" {
				return base + endpoint
			}
			return base
		},
		"params": func(m *parser.Model, r *parser.Route) string {
			return paramTransformer(m, r)
		},
	}
}
