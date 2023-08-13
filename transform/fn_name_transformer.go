package transform

import "strings"

func FnNameTransformer(route string, method string) string {
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
