//strong number or not
//A number can be said as a strong number when the sum of the factorial of the individual digits is equal to the number.For example, 145 is a strong number.
package main

import "fmt"

var sum int = 0
var remainder int = 0
var temp int = 0
var fact int = 0

func main() {
	fmt.Println("enter the number")
	var number int
	fmt.Scanln(&number)
	strongCheck(number)
}

func strongCheck(number int) {
	temp = number
	for number != 0 {
		remainder = number % 10
		fact = factorial(remainder)
		sum += fact
		number = number / 10
	}
	if sum == temp {
		fmt.Println("number is strong")
	} else {
		fmt.Println("number is not strong")
	}
}

func factorial(n int) int {
	var fact int = 1
	for i := 1; i <= n; i++ {
		fact = fact * i
	}
	return fact
}
