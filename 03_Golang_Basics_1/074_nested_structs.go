package main

import "fmt"

type person struct {
	name    string
	age     int
	address address
}

type address struct {
	city  string
	state string
}

func main() {
	a := person{
		name: "John Doe",
		age:  34,
		address: address{
			city:  "Auckland",
			state: "Hamilton",
		},
	}
	fmt.Println("name :", a.name)
	fmt.Println("age : ", a.age)
	fmt.Println("city : ", a.address.city)
	fmt.Println("state : ", a.address.state)
}
