package main

import "fmt"

func changeval(a []int) {
	for i := range a {
		a[i]++
	}
}

func main() {
	b := []int{1, 2, 3, 4, 5}
	fmt.Printf("before function call is %d\n", b)
	changeval(b)
	fmt.Printf("after function call is %d", b)
}
