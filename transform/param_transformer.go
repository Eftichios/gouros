package transform

import (
	"fmt"
	"strings"
	"toncode/gouros/parser"
)

func ParamTransformer(model *parser.Model, route *parser.Route) string {
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
