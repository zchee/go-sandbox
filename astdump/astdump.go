package main

import (
	"flag"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"

	"github.com/davecgh/go-spew/spew"
)

var file = flag.String("f", "", "Dump target go file")

func main() {
	flag.Parse()

	buf, err := ioutil.ReadFile(*file)
	if err != nil {
		log.Fatal(err)
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, *file, buf, parser.Mode(0))
	if err != nil {
		log.Fatal(err)
	}

	s := spew.ConfigState{Indent: "  ", ContinueOnMethod: true}
	s.Dump(f)
}
