package main

import "fmt"

func main() {
	a := [...]string{"a", "b", "c", "d", "e", "f"}
	b := a
	b[0] = "z"
	fmt.Println("a is ", a)
	fmt.Println("b is ", b)
}
