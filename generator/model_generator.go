package generator

import (
	"fmt"
	"os"
	"text/template"
	"toncode/gouros/parser"
)

func GenerateModel(templateResource *parser.TemplateResource) {

	const attributes = "" +
		"{{ define \"attributes\" }}" +
		"{{ range .Model.Attributes }}" +
		"\t{{ .Column }} {{ .Type }}\n" +
		"{{ end }}" +
		"{{ end }}"

	const modelTemplate = "" +
		"package model\n\n" +
		"type {{ template \"structName\" .Model.Table }} struct {\n" +
		"{{ template \"attributes\" . }}" +
		"}"

	funcs := GetFuncs()

	t := template.New("model").Funcs(funcs)
	t.Parse(modelTemplate)
	t.Parse(StructName)
	t.Parse(attributes)

	f, err := os.Create(fmt.Sprintf("model/%s.go", templateResource.Model.Table))

	err = t.Execute(f, templateResource)
	if err != nil {
		panic(err)
	}

}
