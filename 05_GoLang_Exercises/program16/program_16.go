//Sum of numbers in a given range
package main

import "fmt"

func main() {
	fmt.Println("enter min")
	var min int
	fmt.Scanln(&min)
	fmt.Println("enter max")
	var max int
	fmt.Scanln(&max)
	answer := rangeSum(min, max)
	fmt.Println(answer)
}

func rangeSum(min, max int) (sum int) {
	sum = 0
	for i := min; i <= max; i++ {
		sum += i
	}
	return
}
