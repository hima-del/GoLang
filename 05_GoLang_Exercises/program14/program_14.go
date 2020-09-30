//Write program to check if number is power of 3 of any number
package main

import "fmt"

func main() {
	fmt.Println("enter a number")
	var number int
	fmt.Scanln(&number)
	if check(number) == true {
		fmt.Println(" number is power of 3 ")
	} else {
		fmt.Println(" number is not power of 3 of  ")
	}
}

func check(n int) bool {
	for i := 1; i < n/2; i++ {
		if i*i*i == n {
			return true
		}
	}
	return false
}
