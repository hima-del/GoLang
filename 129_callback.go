package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6}
	y := sum(xi...)
	fmt.Println(y)
}

func sum(x ...int) (total int) {
	total = 0
	fmt.Printf("%T\n", x)
	for _, v := range x {
		total += v
	}
	return
}
