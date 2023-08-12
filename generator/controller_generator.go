package generator

import (
	"fmt"
	"os"
	"text/template"
	"toncode/gouros/parser"
)

func GenerateController(templateResource *parser.TemplateResource) {

	const imports = "" +
		"{{ define \"imports\" }}" +
		"import (\n" +
		"\t\"{{ .Module }}/model\"\n" +
		"\t\"{{ .Module }}/service\"\n" +
		"\t\"github.com/gin-gonic/gin\"\n" +
		"\t\"net/http\"\n" +
		")\n" +
		"{{ end }}"

	const controllerAbbr = "" +
		"{{ define \"controllerAbbr\" }}" +
		"{{ . | abbrContr }}" +
		"{{ end }}"

	const serviceAbbr = "" +
		"{{ define \"serviceAbbr\" }}" +
		"{{ . | abbrServ | capitalize }}" +
		"{{ end }}"

	const controllerName = "" +
		"{{ define \"controllerName\" }}" +
		"{{ template \"structName\" . }}Controller" +
		"{{ end }}"

	const serviceName = "" +
		"{{ define \"serviceName\" }}" +
		"{{ template \"structName\" . }}Service" +
		"{{ end }}"

	const methods = "" +
		"{{ define \"methods\" }}" +
		"{{ range .Entity.Routes }}" +
		"func ({{ template \"controllerAbbr\" $.Entity.Resource }} *{{ template \"controllerName\" $.Model.Table }}) " +
		"{{ fnName .Endpoint .Method }}() {}\n" +
		"{{ end }}" +
		"{{ end }}"

	const structDefinition = "" +
		"{{ define \"structDefinition\"}}" +
		"type {{template \"controllerName\" .Entity.Resource}} struct {\n" +
		"\t{{ template \"serviceAbbr\" .Entity.Resource }} *repository.{{ template \"serviceName\" .Entity.Resource}}\n" +
		"}\n" +
		"{{ end }}"

	const controllerTemplate = "" +
		"package controller\n\n" +
		"{{ template \"imports\" . }}\n" +
		"{{ template \"structDefinition\" . }}\n" +
		"{{ template \"methods\" . }}"

	funcs := GetFuncs()
	t := template.New("controller").Funcs(funcs)
	t.Parse(controllerTemplate)
	t.Parse(imports)
	t.Parse(structDefinition)
	t.Parse(methods)
	t.Parse(controllerAbbr)
	t.Parse(controllerName)
	t.Parse(serviceAbbr)
	t.Parse(serviceName)
	t.Parse(StructName)

	f, err := os.Create(fmt.Sprintf("controller/%s_controller.go",
		templateResource.Model.Table))

	err = t.Execute(f, templateResource)
	if err != nil {
		panic(err)
	}

}
