//Find all Factors of a number
package main

import "fmt"

func main() {
	fmt.Println("enter the number")
	var number int
	fmt.Scanln(&number)
	answer := factor(number)
	fmt.Println(answer)
}

func factor(n int) []int {
	var a []int
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			a = append(a, i)
		}
	}
	return a
}
