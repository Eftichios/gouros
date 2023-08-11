package main

import (
	"flag"
	"fmt"
	"toncode/gouros/parser"
)

func main() {

	flag.Parse()

	filename := flag.Arg(0)

    apiResource := parser.ParseYML(filename)

    fmt.Println(apiResource.Entity.Resource)

}
