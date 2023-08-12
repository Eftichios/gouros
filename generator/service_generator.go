package generator

import (
	"fmt"
	"os"
	"text/template"
	"toncode/gouros/parser"
)

func GenerateService(templateResource *parser.TemplateResource) {

	const imports = "" +
		"{{ define \"imports\" }}" +
		"import (\n" +
		"\t\"{{ .Module }}/model\"\n" +
		"\t\"{{ .Module }}/repository\"\n" +
		")\n" +
		"{{ end }}"

	const serviceAbbr = "" +
		"{{ define \"serviceAbbr\" }}" +
		"{{ . | abbrServ }}" +
		"{{ end }}"

	const repoAbbr = "" +
		"{{ define \"repoAbbr\" }}" +
		"{{ . | abbrRepo | capitalize }}" +
		"{{ end }}"

	const serviceName = "" +
		"{{ define \"serviceName\" }}" +
		"{{ template \"structName\" . }}Service" +
		"{{ end }}"

	const repoName = "" +
		"{{ define \"repoName\" }}" +
		"{{ template \"structName\" . }}Repository" +
		"{{ end }}"

	const methods = "" +
		"{{ define \"methods\" }}" +
		"{{ range .Entity.Routes }}" +
		"func ({{ template \"serviceAbbr\" $.Entity.Resource }} *{{ template \"serviceName\" $.Model.Table }}) " +
		"{{ fnName .Endpoint .Method }}() " +
		"{{ .Return }} {}\n" +
		"{{ end }}" +
		"{{ end }}"

	const structDefinition = "" +
		"{{ define \"structDefinition\"}}" +
		"type {{template \"serviceName\" .Entity.Resource}} struct {\n" +
		"\t{{ template \"repoAbbr\" .Entity.Resource }} *repository.{{ template \"repoName\" .Entity.Resource}}\n" +
		"}\n" +
		"{{ end }}"

	const serviceTemplate = "" +
		"package repository\n\n" +
		"{{ template \"imports\" . }}\n" +
		"{{ template \"structDefinition\" . }}\n" +
		"{{ template \"methods\" . }}"

	funcs := GetFuncs()
	t := template.New("repository").Funcs(funcs)
	t.Parse(serviceTemplate)
	t.Parse(imports)
	t.Parse(structDefinition)
	t.Parse(methods)
	t.Parse(serviceAbbr)
	t.Parse(serviceName)
	t.Parse(repoAbbr)
	t.Parse(repoName)
	t.Parse(StructName)

	f, err := os.Create(fmt.Sprintf("service/%s_service.go",
		templateResource.Model.Table))

	err = t.Execute(f, templateResource)
	if err != nil {
		panic(err)
	}

}
