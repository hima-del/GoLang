package main

import (
	"html/template"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	uuid "github.com/satori/go.uuid"
)

type user struct {
	UserName string
	Password []byte
	First    string
	Last     string
}

var tpl *template.Template
var dbusers = map[string]user{}      //user id,user
var dbsessions = map[string]string{} //session id,user id

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	u := getUser(req)
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, req *http.Request) {
	u := getUser(req)
	if !alreadyloggedin(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

func signup(w http.ResponseWriter, req *http.Request) {
	if alreadyloggedin(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	//process form submission
	if req.Method == http.MethodPost {
		un := req.FormValue("usrename")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		//u := user{un, p, f, l}
		//username taken ?
		_, ok := dbusers[un]
		if ok == true {
			http.Error(w, "username already taken", http.StatusSeeOther)
			return
		}

		//create session
		SID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: SID.String(),
		}
		http.SetCookie(w, c)
		dbsessions[c.Value] = un

		//strore user in dbusers
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		u := user{un, bs, f, l}
		dbusers[un] = u
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}
