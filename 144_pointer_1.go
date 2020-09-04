package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)
	b := &a
	fmt.Println(b)
	fmt.Println(*b)
	fmt.Println(*&a)
}
