package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5}
	foo(xi...)
}

func foo(x ...int) (sum int) {
	sum = 0
	fmt.Printf("%T\n", x)
	for i, v := range x {
		sum += v
		fmt.Println("for each element at index postition", i, "we are adding ", v, "and total is", sum)
	}
	return
}
