package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("hello")
	}
	f()
	fmt.Printf("%T", f)
}
