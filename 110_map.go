package main

import "fmt"

func main() {
	detail := map[string]int{
		"john":   32,
		"jenna":  45,
		"Doe":    33,
		"aleena": 23,
	}
	v1, ok := detail["john"]
	fmt.Println(ok)
	fmt.Println(v1)
	v2, ok := detail["sara"]
	fmt.Println(ok)
	fmt.Println(v2)
}
