package main

import "fmt"

func foo(x ...int) (sum int) {
	//fmt.Println(x)
	//fmt.Printf("%T", x) //converted to slice of int
	sum = 0
	for i, v := range x {
		sum = sum + v
		fmt.Println("for each element at index position at", i, "we are adding", v, "and we get", sum)
	}
	return
}

func main() {
	foo(1, 2, 3, 4, 5, 6)
}
