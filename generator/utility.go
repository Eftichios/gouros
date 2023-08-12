package generator

import (
	"strings"
	"text/template"
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

func paramTransformer() {
	return
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
        "createRoute": func(base string, resource string) string {
            if resource != "/" {
                return base + resource
            }
            return base
        },
	}
}
