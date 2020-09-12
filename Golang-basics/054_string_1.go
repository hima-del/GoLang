package main

import "fmt"

func printbytes(str string) {
	for i := 0; i < len(str); i++ {
		fmt.Printf("%x ", str[i])
	}
}

func main() {
	name := "Hello World"
	fmt.Printf("String: %s\n", name)
	printbytes(name)
}
