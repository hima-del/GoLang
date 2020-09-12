package main

import "fmt"

func hello() *int {
	i := 5
	return &i
}

func main() {
	d := hello()
	fmt.Println("value of d is", *d)
}
