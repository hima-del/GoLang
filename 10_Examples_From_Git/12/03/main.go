//It is also possible to pass arguments to anonymous functions just like any other function.
package main

import "fmt"

func main() {
	func(n string) {
		fmt.Println("welcome", n)
	}("himaja")
}
