package main

import "fmt"

func multiply() (num int) {
	x := 4
	y := 10
	num = x * y
	return
}

func main() {
	number := multiply()
	fmt.Printf("%d\n", number)
	switch {
	case number < 20:
		fmt.Printf("%d is less than 20\n", number)
		fallthrough
	case number < 50:
		fmt.Printf("%d is less than 50\n", number)
		fallthrough
	case number < 80:
		fmt.Printf("%d is less than 80", number)
	}
}
