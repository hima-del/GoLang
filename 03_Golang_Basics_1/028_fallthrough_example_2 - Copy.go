package main

import "fmt"

func main() {
	number := 100
	switch {
	case number > 20:
		fmt.Printf("%d is greater than 20\n", number)
		fallthrough
	case number > 50:
		fmt.Printf("%d is greater than 50", number)
	}
}
