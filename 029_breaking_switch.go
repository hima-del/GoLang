package main

import "fmt"

func main() {
	number := -5
	switch {
	case number < 20:
		if number < 0 {
			break
		}
		fmt.Printf("%d is less than 20\n", number)
		fallthrough
	case number < 50:
		fmt.Printf("%d is less than 50\n", number)
		fallthrough
	case number < 80:
		fmt.Printf("%d is less than 80", number)
	}
}
