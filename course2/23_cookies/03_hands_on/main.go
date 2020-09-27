package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("my-cookie")
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "my-cookie",
			Value: "0",
			Path:  "/",
		}
	}
	count, err := strconv.Atoi(cookie.Value)
	if err != nil {
		log.Fatalln(err)
	}
	count++
	cookie.Value = strconv.Itoa(count)
	http.SetCookie(w, cookie) //we have to set the cookie as one of its fields are changed
	io.WriteString(w, cookie.Value)
}
