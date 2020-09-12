package main

import "fmt"

func main() {
	// var a [3]int
	// a[0] = 12
	// a[1] = 15
	// a[2] = 20
	//short hand declaration
	a := [3]int{12, 64, 78}
	fmt.Println(a)
	b := [...]int{28, 45, 79, 60, 50}
	fmt.Println(b)
}
