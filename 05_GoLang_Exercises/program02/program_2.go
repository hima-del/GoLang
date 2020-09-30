//GO Program to take user input and addition of two numbers
package main

import "fmt"

func main() {
	fmt.Println("ENTER FIRST NUMBER")
	var first int
	fmt.Scanln(&first)
	fmt.Println("ENTER SECOND NUMBER")
	var second int
	fmt.Scanln(&second)
	fmt.Println(first + second)
}
