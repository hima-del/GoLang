package main

import "fmt"

func add(x, y int) (sum int) {
	sum = x + y
	return
}

func main() {
	fmt.Println(add(2, 3))
}
