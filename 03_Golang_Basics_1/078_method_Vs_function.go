package main

import "fmt"

type employee struct {
	name     string
	currency string
	salary   int
}

func displaysalary(e employee) {
	fmt.Printf("salary of %s is %s%d", e.name, e.currency, e.salary)
}

func main() {
	employee2 := employee{
		name:     "John Doe",
		currency: "$",
		salary:   23000,
	}
	displaysalary(employee2)
}
