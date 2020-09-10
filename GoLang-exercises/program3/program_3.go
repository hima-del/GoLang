//Enter 2 numbers and display the result for the following operations (each operation should be a function)
//add / multiply / divide / modulus / substract / greater than / less than / is equal
package main

import "fmt"

func main() {
	fmt.Println("ENTER FIRST NUMBER")
	var first float64
	fmt.Scanln(&first)
	fmt.Println("ENTER SECOND NUMBER")
	var second float64
	fmt.Scanln(&second)

	sum := add(first, second)
	fmt.Println("Result of addition: ", sum)

	mult := multiply(first, second)
	fmt.Println("Result of multiplication:", mult)

	div := dividion(first, second)
	fmt.Println("Result of divsion:", div)

	// mod := modulus(first, second)
	// fmt.Println("Result of modulus divsion:", mod)

	sub := subtract(first, second)
	fmt.Println("Result of subtraction:", sub)

	less := lessthan(first, second)
	fmt.Println("Result of lessthan:", less)

	great := greaterthan(first, second)
	fmt.Println("Result of greaterthan:", great)

	equ := equal(first, second)
	fmt.Println("Result of equal:", equ)
}

func add(x, y float64) float64 {
	return x + y
}

func multiply(x, y float64) float64 {
	return x * y
}

func dividion(x, y float64) float64 {
	return x / y
}

// func modulus(x, y float64) float64 {
// 	return x % y
// }

func subtract(x, y float64) float64 {
	return x - y
}

func lessthan(x, y float64) bool {
	if x < y {
		return true
	}
	return false
}

func greaterthan(x, y float64) bool {
	if x > y {
		return true
	}
	return false
}

func equal(x, y float64) bool {
	if x == y {
		return true
	}
	return false
}
