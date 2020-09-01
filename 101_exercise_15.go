package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		rem := i % 4
		fmt.Printf("%d\n", rem)
	}
}
