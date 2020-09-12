package main

import "fmt"

type employee struct {
	firstname string
	lastname  string
	age       int
	salary    int
}

func main() {
	employee1 := &employee{
		firstname: "John",
		lastname:  "Doe",
		age:       34,
		salary:    45000,
	}
	fmt.Println("firstname", employee1.firstname)
	fmt.Println("age", employee1.age)
}
