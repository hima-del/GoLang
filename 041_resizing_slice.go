package main

import "fmt"

func main() {
	letterarray := [...]string{"a", "b", "c", "d", "e"}
	letterslice := letterarray[1:4]
	fmt.Printf("length is %d and capacity is %d\n", len(letterslice), cap(letterslice))
	newslice := letterslice[:cap(letterslice)]
	fmt.Printf("length is %d and capacity is %d", len(newslice), cap(newslice))
}
