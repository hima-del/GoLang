package main

import "fmt"

func main() {
	a := 255
	b := &a
	fmt.Println("address of a is ", b)
	fmt.Println("value of a is ", *b)
	*b++
	fmt.Println("new value of a is", a)
}
