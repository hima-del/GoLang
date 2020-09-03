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
	m := map[string]person{
		p1.lastname: p1,
		p2.lastname: p2,
	}
	for key, value := range m {
		fmt.Println(key)
		fmt.Println(value)
		fmt.Println(value.firstname)
		fmt.Println(value.lastname)
		fmt.Println(value.favouriteicecreamflavours)
		for i, v := range value.favouriteicecreamflavours {
			fmt.Println(i, v)
		}
	}
}
