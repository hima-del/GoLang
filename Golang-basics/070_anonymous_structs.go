package main

import "fmt"

func main() {
	employee3 := struct {
		firstname string
		lastname  string
		age       int
		salary    int
	}{
		firstname: "John",
		lastname:  "Doe",
		age:       34,
		salary:    34000,
	}
	fmt.Println("employee 3 ", employee3)
}
