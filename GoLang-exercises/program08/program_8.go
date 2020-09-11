//Find the sum of elements in a given array. Array should be an input to the program
package main

import "fmt"

func main() {
	fmt.Println("enter the array")
	var a [5]int
	for i := 0; i < len(a); i++ {
		fmt.Scanln(&a[i])
	}
	sum := add(a)
	fmt.Println("sum is :", sum)
}

func add(a [5]int) (sum int) {
	sum = 0
	for _, v := range a {
		sum += v
	}
	return
}
