package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-cookie")
	if err == http.ErrNoCookie {
		c = &http.Cookie{
			Name:  "my-cookie",
			Value: "0",
			//Path:  "/",
		}
	}
	count, err := strconv.Atoi(c.Value)
	if err != nil {
		log.Fatalln(err)
	}
	count++
	c.Value = strconv.Itoa(count)
	http.SetCookie(w, c)
	io.WriteString(w, c.Value)
}
