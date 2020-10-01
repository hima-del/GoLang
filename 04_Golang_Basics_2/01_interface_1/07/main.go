package main

import "fmt"

type describer interface {
	describe()
}

type person struct {
	name string
	age  int
}

func (p person) describe() {
	fmt.Printf("%s is %d years old", p.name, p.age)
}

func findType(i interface{}) {
	switch v := i.(type) {
	case describer:
		v.describe()
	default:
		fmt.Printf("unknown type\n")
	}
}

func main() {
	findType("John Doe")
	p := person{
		name: "John Doe",
		age:  78,
	}
	findType(p)
}
