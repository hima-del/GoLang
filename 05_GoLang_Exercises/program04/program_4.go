//factorial of a number
package main

import "fmt"

func main() {
	fmt.Println("enter the number")
	var n int
	fmt.Scanln(&n)
	answer := factorial(n)
	fmt.Println("factorial of given number is ", answer)
}

func factorial(n int) int {
	var fact int = 1
	for i := 1; i <= n; i++ {
		fact = fact * i
	}
	return fact
}
