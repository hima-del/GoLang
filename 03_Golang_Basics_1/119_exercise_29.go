package main

import "fmt"

type person struct {
	firstname                 string
	lastname                  string
	favouriteicecreamflavours []string
}

func main() {
	p1 := person{
		firstname:                 "John",
		lastname:                  "Doe",
		favouriteicecreamflavours: []string{"pistha", "vanilla", "strawberry"},
	}
	p2 := person{
		firstname:                 "Jennepher",
		lastname:                  "Martin",
		favouriteicecreamflavours: []string{"chocolate", "butterscotch"},
	}
	fmt.Println(p1.firstname)
	fmt.Println(p1.lastname)
	for i, v := range p1.favouriteicecreamflavours {
		fmt.Println(i, v)
	}
	fmt.Println(p2.firstname)
	fmt.Println(p2.lastname)
	for i, v := range p2.favouriteicecreamflavours {
		fmt.Println(i, v)
	}
}
