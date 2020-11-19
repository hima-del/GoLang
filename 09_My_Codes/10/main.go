// This program will print the sum of the squares and cubes of the individual digits of a number.
package main

import (
	"fmt"
)

func main() {
	number := 500
	squarechannel := make(chan int)
	cubechannel := make(chan int)
	go calculateSquares(number, squarechannel)
	go calculateCubes(number, cubechannel)
	sqares, cubes := <-squarechannel, <-cubechannel
	fmt.Println("final output", sqares+cubes)
}

func calculateSquares(number int, squarechannel chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squarechannel <- sum
}

func calculateCubes(number int, cubechannel chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubechannel <- sum
}
