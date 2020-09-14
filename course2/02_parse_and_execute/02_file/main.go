package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml") //here tpl is a pointer to the template(a container that holds all the templates)
	if err != nil {
		log.Fatalln(err)
	}

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln("error creating file", err)
	}
	defer nf.Close()

	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
