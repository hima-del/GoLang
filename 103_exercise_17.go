package main

import "fmt"

func main() {
	x := "hello"
	if x == "bello" {
		fmt.Println("bello")
	} else if x == "hello" {
		fmt.Println("hello")
	} else {
		fmt.Println("neither")
	}
}
