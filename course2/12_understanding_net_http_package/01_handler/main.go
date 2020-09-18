package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) serveHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println("any code you want is this func")
}

func main() {

}
