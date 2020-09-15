package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type person struct {
	Name string
	Age  int
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	p1 := person{
		Name: "John Doe",
		Age:  45,
	}
	err := tpl.Execute(os.Stdout, p1)
	if err != nil {
		log.Fatalln(err)
	}
}
