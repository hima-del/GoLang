//program with 3 threads : one is main go routine which launches 2 other goroutines and main function ends and nothing happens here
package main

import (
	"fmt"
)

func main() {
	go foo()
	go bar()
}

func foo() {
	for i := 0; i < 5; i++ {
		fmt.Println("foo:", i)
	}
}

func bar() {
	for i := 0; i < 5; i++ {
		fmt.Println("bar:", i)
	}
}
