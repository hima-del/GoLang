package main

import "fmt"

func findType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("I am a string and my value is %s\n", i)
	case int:
		fmt.Printf("I am int and my value is %d\n", i)
	default:
		fmt.Printf("Unknown type\n")
	}
}

func main() {
	findType("John Doe")
	findType(89)
	findType(45.55)
}
