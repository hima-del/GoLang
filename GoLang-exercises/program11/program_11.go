//Multiple each element of an array with a constant and print the result
package main

import "fmt"

func main() {
	fmt.Println("enter the array")
	var a [5]int
	for i := 0; i < len(a); i++ {
		fmt.Scanln(&a[i])
	}
	fmt.Println("enetr the number")
	var number int
	fmt.Scanln(&number)
	answer := multiple(a, number)
	fmt.Println(answer)
}

func multiple(a [5]int, number int) (product [5]int) {
	for i := 0; i < len(a); i++ {
		product[i] = a[i] * number
	}
	return
}
