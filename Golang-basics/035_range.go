package main

import "fmt"

func main() {
	a := [...]float64{22.4, 33.5, 55, 98, 55.8, 70}
	var sum float64 = 0
	for i, v := range a {
		fmt.Printf("%d th value of a is %.1f\n", i, v)
		sum = sum + v
	}
	fmt.Printf("the sum is %.1f", sum)
}
