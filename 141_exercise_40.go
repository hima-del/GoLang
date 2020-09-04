package main

import "fmt"

func foo() func() int {
	return func() int {
		return 42
	}
}

func main() {
	f := foo()
	g := f()
	//fmt.Println(f)
	fmt.Println(g)
}
