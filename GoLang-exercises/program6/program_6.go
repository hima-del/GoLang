//Program to Check Prime Number
package main

import "fmt"

func main() {
	var n int
	var flag int = 0
	fmt.Println("enter the number")
	fmt.Scanln(&n)
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			flag = 1
		}
	}
	if n == 1 {
		fmt.Println("1 is neither composite nor prime")
	} else {
		if flag == 0 {
			fmt.Println("number is prime")
		} else {
			fmt.Println("number is not prime")
		}
	}
}
