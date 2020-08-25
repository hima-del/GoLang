package main

import "fmt"

func changeVal(num [5]int) {
	num[0] = 11
	fmt.Println("inside the function", num)
}

func main() {
	num := [...]int{6, 9, 5, 4, 3}
	fmt.Println("before function call", num)
	changeVal(num)
	fmt.Println("after function call", num)
}
