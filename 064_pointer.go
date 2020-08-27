package main

import "fmt"

func main() {
	a := 255
	var b *int = &a
	fmt.Println("address of a is", b)
}
