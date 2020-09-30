package main

import "fmt"

func main() {
	x := bar()
	//fmt.Println(x)
	fmt.Printf("%T\n", x)
	y := x()
	fmt.Println(y)
}

func bar() func() int {
	return func() int {
		return 451
	}
}
