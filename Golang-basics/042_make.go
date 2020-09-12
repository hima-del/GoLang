package main

import "fmt"

func main() {
	i := make([]int, 5, 5)
	fmt.Println("slice is", i)
	fmt.Printf("slice is %d", i)
}
