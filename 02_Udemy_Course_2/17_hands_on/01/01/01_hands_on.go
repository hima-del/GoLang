//ListenAndServe on port ":8080" using the default ServeMux.
//Use HandleFunc to add the following routes to the default ServeMux:
//"/" "/dog/" "/me/
//Add a func for each of the routes.
//Have the "/me/" route print out your name.
package main

import (
	"io"
	"net/http"
)

func c(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "dog dog")
}

func d(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello himaja")
}

func main() {
	http.HandleFunc("/dog/", c)
	http.HandleFunc("/me/", d)
	http.ListenAndServe(":8080", nil)
}
