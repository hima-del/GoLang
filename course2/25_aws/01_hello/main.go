package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":80", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "oh yeah, I'm running on AWS")
}
