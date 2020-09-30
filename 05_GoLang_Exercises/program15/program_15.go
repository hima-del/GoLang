//Sum of digits of a number
package main

import "fmt"

func main() {
	fmt.Println("enter the number")
	var number int
	fmt.Scanln(&number)
	answer := addDigit(number)
	fmt.Println(answer)
}

func addDigit(n int) int {
	sum := 0
	remainder := 0
	for n != 0 {
		remainder = n % 10
		sum += remainder
		n = n / 10
	}
	return sum
}
