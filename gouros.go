package main

import (
	"flag"
	"fmt"
	"toncode/gouros/generator"
	"toncode/gouros/parser"
)


func main() {

	flag.Parse()

	filename := flag.Arg(0)

	apiResource := parser.ParseYML(filename)
	module := parser.ParseGoMod()

	fmt.Println(apiResource.Entity.Resource)

	for _, model := range apiResource.Models {

		templateResource := &parser.TemplateResource{
			Model:  model,
			Entity: apiResource.Entity,
			Module: module,
		}

		generator.GenerateModel(templateResource)
        generator.GenerateRepository(templateResource)
        generator.GenerateService(templateResource)
        generator.GenerateController(templateResource)
        generator.GenerateRouter(templateResource)
	}

}
