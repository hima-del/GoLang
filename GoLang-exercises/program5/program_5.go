// Golang program to check whether a number is palindrome or not
package main

import "fmt"

func main() {
	fmt.Println("enter the number to be reversed")
	var remainder, number, temp int
	fmt.Scanln(&number)
	temp = number
	var reverse int = 0

	for {
		remainder = number % 10
		reverse = reverse*10 + remainder
		number = number / 10

		if number == 0 {
			break
		}
	}
	if temp == reverse {
		fmt.Println("number is pallindrome")
	} else {
		fmt.Println("number is not pallindrome")
	}
}
