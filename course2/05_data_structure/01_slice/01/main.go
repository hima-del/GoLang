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

func main() {

	fruits := []string{"pappaya", "orange", "apple", "banana", "mango"}
	err := tpl.Execute(os.Stdout, fruits)
	if err != nil {
		log.Fatalln(err)
	}
}

//{{range .}} renge over the current piece of data here we are using slice
//<li>{{.}}</li>//here the current piece becomes the elements of slice
//{{end}}
