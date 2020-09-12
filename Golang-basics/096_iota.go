package main

import "fmt"

const (
	a = iota
	b
	_
	d
)

func main() {
	fmt.Println(a, b, d)
}
