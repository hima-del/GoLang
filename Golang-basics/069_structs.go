package main

import "fmt"

type employee struct {
	firstname string
	lastname  string
	age       int
	salary    int
}

func main() {
	employee1 := employee{
		firstname: "John",
		lastname:  "Doe",
		salary:    34000,
		age:       43,
	}
	employee2 := employee{"Brad", "Traversy", 34, 39000}
	fmt.Println("employee 1", employee1)
	fmt.Println("employee 2", employee2)
}
