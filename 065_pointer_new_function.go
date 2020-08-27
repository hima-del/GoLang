package main

import "fmt"

func main() {
	size := new(int)
	fmt.Printf("size value is %d,type is %T,address is %v\n", *size, size, size)
	*size = 85
	fmt.Println("new size value is ", *size)
}
