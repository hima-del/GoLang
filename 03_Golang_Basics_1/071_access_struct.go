package main

import "fmt"

type employee struct {
	firstname string
	lastname  string
	age       int
	salary    int
}

func main() {
	employee4 := employee{
		firstname: "John",
		lastname:  "Doe",
		age:       34,
		salary:    56000,
	}
	fmt.Println("firstname is ", employee4.firstname)
	fmt.Println("lastname is ", employee4.lastname)
	fmt.Printf("age is %d \n", employee4.age)
	fmt.Printf("salary is %d\n ", employee4.salary)
	employee4.salary = 77000
	fmt.Printf("new salary is %d", employee4.salary)
}
