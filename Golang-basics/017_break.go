package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		if i > 5 {
			break
		}
		fmt.Printf("%d\n", i)
	}
	fmt.Printf("Line after for loop")
}
