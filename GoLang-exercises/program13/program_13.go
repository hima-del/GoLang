// Write a program that asks the user for a number n and gives them the possibility to choose between computing the sum and computing the product of 1,…,n.
package main

import "fmt"

func main() {
	fmt.Println("eneter your choice : product or sum")
	var choice string
	fmt.Scanln(&choice)
	fmt.Println("enter a number")
	var number int
	fmt.Scanln(&number)
	if choice == "product" {
		productanswer := product(number)
		fmt.Println("result of product : ", productanswer)
	} else {
		sumanswer := add(number)
		fmt.Println("result of sum : ", sumanswer)
	}
}

func product(n int) int {
	var pro int = 1
	for i := 1; i <= n; i++ {
		pro *= i
	}
	return pro
}

func add(n int) int {
	var sum int = 0
	for i := 0; i <= n; i++ {
		sum += i
	}
	return sum
}
