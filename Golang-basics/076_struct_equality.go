package main

import "fmt"

type person struct {
	firstname string
	lastname  string
}

func main() {
	person1 := person{
		firstname: "John",
		lastname:  "Doe",
	}
	person2 := person{
		firstname: "John",
		lastname:  "Doe",
	}
	if person1 == person2 {
		fmt.Println("person 1 and person 2 are equal")
	} else {
		fmt.Println("person 1 and person 2 are not equal")
	}
	person3 := person{
		firstname: "Anna",
	}
	person4 := person{
		firstname: "Anna",
		lastname:  "Xaviour",
	}
	if person3 == person4 {
		fmt.Println("person 3 and person 4 are equal")
	} else {
		fmt.Println("person 3 and person 4 are not equal")
	}
}
