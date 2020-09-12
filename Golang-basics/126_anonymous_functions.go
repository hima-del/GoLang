//does not have name
package main

import "fmt"

func main() {
	func(x int) {
		fmt.Println("anonymous function", x)
	}(42)
}
