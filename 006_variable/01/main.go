package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

type exampleStruct struct {
	Name string
	Age  int
}

func main() {

	persion := exampleStruct{"thanh", 24}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", persion)
	if err != nil {
		log.Fatalln(err)
	}
}
