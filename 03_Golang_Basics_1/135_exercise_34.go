package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := foo(a...)
	fmt.Println(b)
	c := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Println(bar(c))
}

func foo(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func bar(x []int) int {
	fmt.Printf("%T\n", x)
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
