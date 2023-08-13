package generator

import (
	"fmt"
	"os"
	"text/template"
	"toncode/gouros/parser"
)

func GenerateRouter(templateResource *parser.TemplateResource) {

	const imports = "" +
		"{{ define \"imports\" }}" +
		"import (\n" +
		"\t\"database/sql\"\n" +
		"\t\"{{ .Module }}/controller\"\n" +
		"\t\"{{ .Module }}/repository\"\n" +
		"\t\"{{ .Module }}/service\"\n" +
		"\t\"github.com/gin-gonic/gin\"\n" +
		")\n" +
		"{{ end }}"

	const repoAbbr = "" +
		"{{ define \"repoAbbr\" }}" +
		"{{ . | abbrRepo }}" +
		"{{ end }}"

	const repoName = "" +
		"{{ define \"repoName\" }}" +
		"{{ template \"structName\" . }}Repository" +
		"{{ end }}"

	const serviceAbbr = "" +
		"{{ define \"serviceAbbr\" }}" +
		"{{ . | abbrServ }}" +
		"{{ end }}"

	const serviceName = "" +
		"{{ define \"serviceName\" }}" +
		"{{ template \"structName\" . }}Service" +
		"{{ end }}"

	const controllerAbbr = "" +
		"{{ define \"controllerAbbr\" }}" +
		"{{ . | abbrContr }}" +
		"{{ end }}"

	const controllerName = "" +
		"{{ define \"controllerName\" }}" +
		"{{ template \"structName\" . }}Controller" +
		"{{ end }}"

	const init = "" +
		"{{ define \"init\" }}" +
		"\t{{ template \"repoAbbr\" . }} := " +
		"&repository.{{ template \"repoName\" . }}{Db: db}\n" +
		"\t{{ template \"serviceAbbr\" . }} := " +
		"&service.{{ template \"serviceName\" . }}{ " +
		"{{ . | abbrRepo | capitalize }}: {{ template \"repoAbbr\" . }} }\n" +
		"\t{{ template \"controllerAbbr\" . }} := " +
		"&controller.{{ template \"controllerName\" . }}{ " +
		"{{ . | abbrServ | capitalize }}: {{ template \"serviceAbbr\" . }} }\n" +
		"{{ end }}"

	const routes = "" +
		"{{ define \"routes\"}}" +
		"{{ range .Entity.Routes }}" +
		"\tgroup.{{ .Method }}(\"{{ createRoute $.Entity.Base .Endpoint }}\", " +
		"{{ template \"controllerAbbr\" $.Entity.Resource }}.{{ fnName .Endpoint .Method  }})\n" +
		"{{ end }}" +
		"{{ end }}"

	const routerTemplate = "" +
		"package router\n\n" +
		"{{ template \"imports\" . }}" +
		"func New{{ template \"structName\" .Entity.Resource }}Router(" +
		"db *sql.DB, group *gin.RouterGroup) {\n" +
		"{{ template \"init\" .Entity.Resource}}" +
		"{{ template \"routes\" .}}" +
		"}"

	funcs := GetFuncs()
	t := template.New("router").Funcs(funcs)
	t.Parse(routerTemplate)
	t.Parse(imports)
	t.Parse(init)
	t.Parse(routes)
	t.Parse(repoAbbr)
	t.Parse(repoName)
	t.Parse(controllerAbbr)
	t.Parse(controllerName)
	t.Parse(serviceAbbr)
	t.Parse(serviceName)
	t.Parse(StructName)

	f, err := os.Create(fmt.Sprintf("router/%s_router.go",
		templateResource.Model.Table))

	err = t.Execute(f, templateResource)
	if err != nil {
		panic(err)
	}
}
