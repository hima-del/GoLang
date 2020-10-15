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

func init(){
	var err error
	db,err:=sql.Open("postgres","postgres://bond:password@localhost/bookstore?sslmode=disable")
	if err!=nil{
		panic(err)
	}
	err=db.Ping()
	if err!=nil{
		panic(err)
	}
	fmt.Println("you are connected to database")

	tpl=template.Must(template.ParseGlob("templates/*.gohtml"))
}

type Book struct{
	Isbn string
	Title string
	Author string
	Price float32
}

func main(){
	http.HandleFunc("/",index)
	http.HandleFunc("/books", booksIndex)
	http.HandleFunc("/books/show", booksShow)
	http.HandleFunc("/books/create", booksCreateForm)
	http.HandleFunc("/books/create/process", booksCreateProcess)
	http.ListenAndServe(":8080", nil)
}

func index(w http.)



