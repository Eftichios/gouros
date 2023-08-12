package generator

import (
	"fmt"
	"os"
	"text/template"
	"toncode/gouros/parser"
)

func GenerateRepository(templateResource *parser.TemplateResource) {

	const imports = "" +
		"{{ define \"imports\" }}" +
		"import (\n" +
		"\t\"database/sql\"\n" +
		"\t\"{{ .Module }}/model\"\n" +
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

	const methods = "" +
		"{{ define \"methods\" }}" +
		"{{ range .Entity.Routes }}" +
		"func ({{ template \"repoAbbr\" $.Entity.Resource }} *{{ template \"repoName\" $.Model.Table }}) " +
		"{{ fnName .Endpoint .Method }}() " +
		"{{ .Return }} {}\n" +
		"{{ end }}" +
		"{{ end }}"

	const structDefinition = "" +
		"{{ define \"structDefinition\"}}" +
		"type {{template \"repoName\" .Entity.Resource}} struct {\n" +
		"\tDb *sql.DB\n" +
		"}\n" +
		"{{ end }}"

	const repositoryTemplate = "" +
		"package repository\n\n" +
		"{{ template \"imports\" . }}\n" +
		"{{ template \"structDefinition\" . }}\n" +
		"{{ template \"methods\" . }}"

	funcs := GetFuncs()
	t := template.New("repository").Funcs(funcs)
	t.Parse(repositoryTemplate)
	t.Parse(imports)
	t.Parse(structDefinition)
	t.Parse(methods)
	t.Parse(repoAbbr)
	t.Parse(repoName)
	t.Parse(StructName)

	f, err := os.Create(fmt.Sprintf("repository/%s_repository.go",
		templateResource.Model.Table))

	err = t.Execute(f, templateResource)
	if err != nil {
		panic(err)
	}

}
