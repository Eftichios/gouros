package generator

import (
	"fmt"
	"os"
	"text/template"
	"toncode/gouros/parser"
)

func GenerateModel(model parser.Model) {

	const structName = "" +
		"{{ define \"structName\" }}" +
		"{{ .Table | capitalize }}" +
		"{{ end -}}"

	const attributes = "" +
		"{{ define \"attributes\" }}" +
		"{{ range .Attributes }}" +
		"\t{{ .Column }} {{ .Type }}\n" +
		"{{ end }}" +
		"{{ end }}"

	const modelTemplate = "" +
		"package model\n\n" +
		"type {{ template \"structName\" . }} struct {\n" +
		"{{ template \"attributes\" . }}" +
		"}"

	funcs := GetFuncs()

	t := template.New("model").Funcs(funcs)
	t.Parse(modelTemplate)
	t.Parse(structName)
	t.Parse(attributes)

	f, err := os.Create(fmt.Sprintf("model/%s.go", model.Table))

	err = t.Execute(f, model)
	if err != nil {
		panic(err)
	}

}
