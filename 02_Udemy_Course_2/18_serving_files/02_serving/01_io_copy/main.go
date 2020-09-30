package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/toby.jpg", dogPic)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `
	<img src="/toby.jpg">
	`)
}

func dogPic(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("toby.jpg") //here we get pointer to a file and it implements a reader and writer interface ie,we can write to it and read from it
	if err != nil {
		http.Error(w, "filr not found", 404)
		return
	}
	defer f.Close()
	io.Copy(w, f) //we are copying from the file to the response writer
}
