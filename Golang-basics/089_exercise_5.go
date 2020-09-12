package main

import "fmt"

type hotdog int

var x hotdog
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\t", x)
	x = 10
	fmt.Println(x)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T", y)
}
