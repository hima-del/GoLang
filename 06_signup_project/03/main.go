package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

var db *sql.DB
var tpl *template.Template

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://himaja:password@localhost/project?sslmode=disable")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("you are connected to database.")
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

type User struct {
	Username string
	Password string
}

func main() {
	http.HandleFunc("/users/create", usersCreateForm)
	http.HandleFunc("/users/create/process", usersCreateProcess)
	http.ListenAndServe(":8080", nil)
}

func usersCreateForm(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "create.gohtml", nil)
}

func usersCreateProcess(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	//get form values
	u := User{}
	u.Username = req.FormValue("username")
	u.Password = req.FormValue("password")

	//validate form values
	if u.Username == "" || u.Password == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	//insert values
	_, err := db.Exec("insert into users (username,password)VALUES($1,$2)", u.Username, u.Password)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	//confirm insertion
	tpl.ExecuteTemplate(w, "created.gohtml", u)
}
