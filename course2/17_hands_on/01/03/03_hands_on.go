//Take the previous program and change it so that:
//func main uses http.Handle instead of http.HandleFunc
//Contstraint: Do not change anything outside of func main
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
	http.Handle("/dog/", http.HandlerFunc(c))
	http.Handle("/cat/", http.HandlerFunc(d))
	http.Handle("/me/", http.HandlerFunc(name))
	http.ListenAndServe(":8080", nil)
}
