package main

import "fmt"

type employee struct {
	name string
	age  int
}

//method with value reciever

func (e employee) changename(newname string) {
	e.name = newname
}

//method with pointer reciever

func (e *employee) changeage(newage int) {
	e.age = newage
}

func main() {
	employee1 := employee{
		name: "John Doe",
		age:  34,
	}
	fmt.Printf("name before function call is %s\n", employee1.name)
	employee1.changename("John Smith")
	fmt.Printf("name after function call is %s\n", employee1.name)
	fmt.Printf("age before function call is %d\n", employee1.age)
	employee1.changeage(66)
	fmt.Printf("age after function call is %d\n", employee1.age)
}
