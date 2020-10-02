//passing functions as arguments to other functions
package main

import "fmt"

func simple(a func(a, b int) int) {
	fmt.Println(a(7, 8))
}

func main() {
	f := func(a, b int) int {
		return a + b
	}
	simple(f)
}
