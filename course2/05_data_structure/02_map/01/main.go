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

	fruits := map[string]string{
		"one":   "apple",
		"two":   "orange",
		"three": "mango",
		"four":  "banana",
	}
	err := tpl.Execute(os.Stdout, fruits)
	if err != nil {
		log.Fatalln(err)
	}
}
