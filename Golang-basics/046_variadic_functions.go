package main

import "fmt"

func findinteger(number int, numberarray ...int) {
	found := false
	for i, v := range numberarray {
		if number == v {
			fmt.Printf("number %d is found at position %d of array %d\n", number, i, numberarray)
			found = true
		}
	}
	if found == false {
		fmt.Printf("number not found\n")
	}
}

func main() {
	findinteger(2, 4, 5, 6, 2)
	findinteger(1, 2, 3, 4, 5, 6, 7)
}
