package main

import "fmt"

type employee struct {
	string
	int
}

func main() {
	employee2 := employee{
		string: "John Doe",
		int:    34,
	}
	fmt.Println("name is ", employee2.string)
	fmt.Println("age is ", employee2.int)
}
