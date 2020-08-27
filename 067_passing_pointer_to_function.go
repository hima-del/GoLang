package main

import "fmt"

func changeval(a *int) {
	*a = 55
}

func main() {
	a := 50
	b := &a
	fmt.Println("value of a before function call", a)
	changeval(b)
	fmt.Println("value of a after function call", a)
}
