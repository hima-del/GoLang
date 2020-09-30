//Write a program that asks the user for a number n and prints the sum of the numbers 1 to n
package main

import "fmt"

func main() {
	fmt.Println("enter the number")
	var number int
	fmt.Scanln(&number)
	sum := add(number)
	fmt.Println("Sum of numbers upto", number, "is", sum)
}

func add(n int) (sum int) {
	sum = 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return
}
