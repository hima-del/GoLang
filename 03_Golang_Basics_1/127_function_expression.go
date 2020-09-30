package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("Hello world")
	}
	f()
	g := func(x int) {
		fmt.Println("hi, there are ", x, "roses")
	}
	g(10)
}
