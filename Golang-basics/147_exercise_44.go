package main

import "fmt"

type person struct {
	name string
	age  int
}

func changeme(p *person) {
	p.name = "John Doe"
}

func main() {
	p1 := person{
		name: "James Bond",
		age:  45,
	}
	fmt.Println(p1)
	changeme(&p1)
	fmt.Println(p1)
}
