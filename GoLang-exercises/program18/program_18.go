//Armstrong number or not
package main

import "fmt"

var sum int = 0
var remainder int = 0
var temp int = 0

func main() {
	fmt.Println("enter the number")
	var number int
	fmt.Scanln(&number)
	checkamstrong(number)
}

func checkamstrong(number int) {
	temp = number
	for number != 0 {
		remainder = number % 10
		sum += remainder * remainder * remainder
		number = number / 10
	}
	if temp == sum {
		fmt.Println("number is amstrong")
	} else {
		fmt.Println("number is not amstrong")
	}
}
