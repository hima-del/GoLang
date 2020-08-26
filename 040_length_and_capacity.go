package main

import "fmt"

func main() {
	a := [...]string{"a", "b", "c", "d", "e", "f"}
	var b []string = a[1:4]
	fmt.Printf("length is %d and capacity is %d", len(b), cap(b))
}
