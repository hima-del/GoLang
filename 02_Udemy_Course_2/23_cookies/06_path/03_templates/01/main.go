package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/bowser", bowser)
	http.HandleFunc("/dog/bowser/pictures", bowserpics)
	http.HandleFunc("/cat", cat)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	var c *http.Cookie
	c, err := req.Cookie("user-cookie")
	if err != nil {
		fmt.Println(err)
		fmt.Printf("%T\n", c)
	}
	tpl.ExecuteTemplate(w, "index.gohtml", c)
}

func bowser(w http.ResponseWriter, req *http.Request) {
	c := &http.Cookie{
		Name:  "user-cookie",
		Value: "this would be the value",
		Path:  "/dog/bowser",
	}
	http.SetCookie(w, c)
	tpl.ExecuteTemplate(w, "browser.gohtml", c)
}

func bowserpics(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("user-cookie")
	if err != nil {
		fmt.Println(err)
		fmt.Printf("%T\n", c)
	}
	tpl.ExecuteTemplate(w, "browserpics.gohtml", c)
}

func cat(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("user-cookie")
	if err != nil {
		fmt.Println(err)
		fmt.Printf("%T\n", c)
	}
	tpl.ExecuteTemplate(w, "cat.gohtml", c)
}
