package main

import "fmt"

var x int

func main() {
	fmt.Println(x)
	x++
	{
		y := 100
		fmt.Println(y)
	}
	//fmt.Println(y)
	fmt.Println(x)
	foo()
	fmt.Println(x)
}

func foo() {
	fmt.Println("hello")
	x++
}
