package main

import "fmt"

func main() {
	s := struct {
		first     string
		friends   map[string]int
		favdrinks []string
	}{
		first: "John",
		friends: map[string]int{
			"Doe":    33,
			"Anniee": 30,
		},
		favdrinks: []string{"wtaer", "lemon juice", "cola"},
	}
	fmt.Println(s)
	fmt.Println(s.first)
	fmt.Println(s.friends)
	fmt.Println(s.favdrinks)
	for key, value := range s.friends {
		fmt.Println(key, value)
		//fmt.Println(value)
	}
	for i, v := range s.favdrinks {
		fmt.Println(i, v)
	}
}
