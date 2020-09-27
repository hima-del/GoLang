package main

import (
	"github.com/satori/go.uuid"
	"net/http"
	"html/template"
)

type user struct{
	UserName string
	First string
	Last string 
}

var tpl *template.Template
var dbUsers=map[string]user //user id,user
var dbSessions =map[string]string //session id,user id

func init(){
	tpl=template.Must(template.ParseGlob("templates/*"))
}

func main(){
	http.HandleFunc("/",foo)
	http.HandleFunc("/bar",bar)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":8080",nil)
}

func foo(w http.ResponseWriter,req *http.Request){
	c,err:=req.Cookie("session")
	if err!=nil{
		SID,_:=uuid.NewV4()
		c=&http.Cookie{
			Name:"session",
			Value:SID.String(),
		}
		http.SetCookie(w,c)
	}

	//if the user already exists, get the user
	var u user
	un,ok:=dbSessions[c.Value]
	if ok==true{
		u=dbUsers[un]
	}

	//form submission

	if req.Method==http.MethodPost{
		
	}
}