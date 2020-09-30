package main

import "fmt"

func main() {
	var letter string = "i"
	fmt.Printf("the letter %s is a ", letter)
	switch letter {
	case "a", "e", "i", "o", "u":
		fmt.Println("vowel")
	default:
		fmt.Println("nit a vowel")
	}
}
