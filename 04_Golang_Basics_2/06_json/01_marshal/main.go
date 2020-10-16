package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "John",
		Last:  "Doe",
		Age:   44,
	}

	p2 := person{
		First: "anna",
		Last:  "jackson",
		Age:   34,
	}
	people := []person{p1, p2}

	//fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
