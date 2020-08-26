package main

import "fmt"

func printarray(a [2][3]string) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}

func main() {
	var b [2][3]string
	b[0][0] = "cat"
	b[0][1] = "dog"
	b[0][2] = "cow"
	b[1][0] = "rat"
	b[1][1] = "elephant"
	b[1][2] = "deer"
	printarray(b)
	c := [2][3]string{
		{"a", "b", "c"},
		{"d", "e", "f"},
	}
	printarray(c)
}
