package main

import "fmt"

func foo(x ...int) {
	fmt.Println(x)
	fmt.Printf("%T", x) //converted to slice of int
}

func main() {
	foo(1, 2, 3, 4, 5, 6)
}
