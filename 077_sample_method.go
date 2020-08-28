package main

import "fmt"

type employee struct {
	name     string
	salary   int
	currency string
}

func (e employee) displaysalary() {
	fmt.Printf("salary of %s is %s%d", e.name, e.currency, e.salary)
}

func main() {
	employee1 := employee{
		name:     "John Doe",
		salary:   34000,
		currency: "$",
	}
	employee1.displaysalary()
}
