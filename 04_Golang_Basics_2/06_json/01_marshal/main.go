package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "John",
		last:  "Doe",
		age:   44,
	}

	p2 := person{
		first: "anna",
		last:  "jackson",
		age:   34,
	}
	people := []person{p1, p2}

	fmt.Println(people)
}
