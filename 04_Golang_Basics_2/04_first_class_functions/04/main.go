//user defined function types
package main

import "fmt"

type add func(a, b int) int

func main() {
	var a add = func(a, b int) int {
		return a + b
	}
	s := a(5, 6)
	fmt.Println(s)
}
