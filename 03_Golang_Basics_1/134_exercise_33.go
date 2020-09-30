package main

import "fmt"

func main() {
	n := foo()
	x, s := bar()
	fmt.Println(n, x, s)
}

func foo() int {
	return 67
}

func bar() (int, string) {
	return 55, "hello world"
}
