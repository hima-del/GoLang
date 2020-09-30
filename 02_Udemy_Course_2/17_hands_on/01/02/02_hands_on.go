//Take the previous program in the previous folder and change it so that:
//a template is parsed and served
//you pass data into the template
package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

var tpl *template.Template

func c(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "dog dog")
}

func d(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello himaja")
}
func name(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(w, "tpl.gohtml", "HIMAJA")
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	http.HandleFunc("/dog/", c)
	http.HandleFunc("/cat/", d)
	http.HandleFunc("/me/", name)
	http.ListenAndServe(":8080", nil)
}
