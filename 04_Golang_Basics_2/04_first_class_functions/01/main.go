//anonymous functions
package main

import "fmt"

func main() {
	a := func() {
		fmt.Println("hello world first class function")
	}
	a()
	fmt.Printf("%T", a)
}

//the function assigned to a does not have a name.
//These kind of functions are called anonymous functions since they do not have a name.
