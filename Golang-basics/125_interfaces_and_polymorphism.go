package main

import "fmt"

//keyword identifier type
//var x int
//type person struct
//type human interface
//a value can be of more than one type
type person struct {
	firstname string
	lastname  string
}

type secretagent struct {
	person
	ltk bool
}

type human interface { //any type that has a method speak() is also of type human
	speak()
}

func bar(h human) {
	fmt.Println("i was passed into bar", h)
}

func (s secretagent) speak() { //this function can be attached to anybody that has a type secretagent
	fmt.Println("hi i am", s.firstname, s.lastname, "-the secret agent speak")
}

func (p person) speak() { //this function can be attached to anybody that has a type secretagent
	fmt.Println("hi i am", p.firstname, p.lastname, "-the person speak")
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
	p1 := person{
		firstname: "Dr",
		lastname:  "Yes",
	}
	fmt.Println(sa1)
	sa1.speak()
	fmt.Println(sa2)
	sa2.speak()
	fmt.Println(p1)
	bar(sa1)
	bar(sa2)
	bar(p1)
}
