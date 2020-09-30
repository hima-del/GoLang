// Write a program that asks the user for a number n and prints the sum of the numbers 1 to n such that only multiples of three or five are considered in the sum, e.g. 3, 5, 6, 9, 10, 12, 15 for n=17
package main

import "fmt"

func main() {
	fmt.Println("enter the number")
	var number int
	fmt.Scanln(&number)
	answer := add(number)
	fmt.Println(answer)
}

func add(n int) (sum int) {
	sum = 0
	for i := 0; i <= n; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return
}
