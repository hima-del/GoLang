package main

import "fmt"

type person struct {
	name string
	age  int
	address
}

type address struct {
	city  string
	state string
}

func (a address) fulladdress() {
	fmt.Printf("full address is %s%s", a.city, a.state)
}

func main() {
	person1 := person{
		name: "John Doe",
		age:  34,
		address: address{
			city:  "Auckland",
			state: "Hamilton",
		},
	}
	person1.fulladdress()

}
