//It is also possible to call a anonymous function without assigning it to a variable.
package main

import "fmt"

func main() {
	func() {
		fmt.Println("hello world first class function")
	}()
}

//we call the function using ()
