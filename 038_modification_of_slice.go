package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var b []int = a[1:6]
	fmt.Printf("array before %d\n", a)
	for i := range b {
		b[i]++
	}
	fmt.Printf("array after %d", a)
}
