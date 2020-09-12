package main

import "fmt"

func main() {
	p1 := struct {
		firstname string
		lastname  string
		age       int
	}{
		firstname: "John",
		lastname:  "Doe",
		age:       34,
	}
	fmt.Println(p1)
}
