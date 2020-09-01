package main

import "fmt"

type value int

var x value

func main() {
	fmt.Println(x)
	// fmt.Printf("%d\n", x)
	fmt.Printf("%T\t", x)
	x = 42
	fmt.Println(x)
}
