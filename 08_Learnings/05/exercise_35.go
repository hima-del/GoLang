package main

import "fmt"

func main() {
	defer foo("hello", "world")
	//fmt.Println(a)
	bar(100, 10)
	//fmt.Println(b)
}

func foo(x string, y string) string {
	z := x + " " + y
	fmt.Println(z)
	return z
}

func bar(x int, y int) int {
	z := x / y
	fmt.Println(z)
	return z
}
