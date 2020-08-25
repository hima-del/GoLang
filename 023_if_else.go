package main

import "fmt"

func main() {
	num := 11
	if num%2 == 0 {
		fmt.Println("the number", num, "is even")
	} else {
		fmt.Println("the number", num, "is odd")
	}
}
