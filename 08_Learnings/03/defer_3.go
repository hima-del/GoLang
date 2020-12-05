package main

import "fmt"

func main() {
	a := c()
	fmt.Println(a)
}

func c() (i int) {
	defer func() {
		i++
	}()
	return 1
}
