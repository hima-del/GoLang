package main

import "fmt"

func simple() func(a, b int) int {
	f := func(a, b int) int {
		return a + b
	}
	return f
}

func main() {
	s := simple()
	fmt.Println(s(30, 60))
}

// The return value from simple is assigned to s.
//Now s contains the function returned by simple function.
// We call s and pass it two int arguments
