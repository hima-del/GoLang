package main

import "fmt"

type person struct {
	firstname string
	lastname  string
}

type secretagent struct {
	person
	ltk bool
}

func (s secretagent) speak() { //this function can be attached to anybody that has a type secretagent
	fmt.Println("hi i am", s.firstname, s.lastname)
}

func main() {
	sa1 := secretagent{
		ltk: true,
		person: person{
			firstname: "John",
			lastname:  "Doe",
		},
	}
	sa2 := secretagent{
		ltk: false,
		person: person{
			firstname: "Anna",
			lastname:  "Xaviour",
		},
	}
	fmt.Println(sa1)
	sa1.speak()
	fmt.Println(sa2)
	sa2.speak()
}
