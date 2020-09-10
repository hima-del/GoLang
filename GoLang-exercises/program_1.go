//GO Program to take user input and addition of two strings
package main

import "fmt"

func main() {
	fmt.Println("enter the first string")
	var first string
	fmt.Scanln(&first)
	fmt.Println("enter the second string")
	var second string
	fmt.Scanln(&second)
	fmt.Println(first + second)
}
