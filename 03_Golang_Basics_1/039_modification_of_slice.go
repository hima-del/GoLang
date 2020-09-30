package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var b []int = a[:]
	var c []int = a[:]
	fmt.Printf("array is %d\n", a)
	b[0] = 10
	fmt.Printf("array is %d\n", a)
	c[1] = 20
	fmt.Printf("array is %d", a)
}
