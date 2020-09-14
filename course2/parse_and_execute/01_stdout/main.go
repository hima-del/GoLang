////understanding the package text/template
//parsing and executing templates
package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.Execute(os.Stdout, nil) //executes the template and writes to new file
	if err != nil {
		log.Fatalln(err)
	}
}
//once we have parsed we have a `pointer to template`(container that holds all the templates) 
//when we have pointer to template we can execute
