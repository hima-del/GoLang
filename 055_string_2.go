package main

import "fmt"

func printbytes(str string) {
	fmt.Printf("bytes:")
	for i := 0; i < len(str); i++ {
		fmt.Printf("%x ", str[i])
	}
}

func printcharacter(str string) {
	fmt.Printf("characters:")
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c ", str[i])
	}
}

func main() {
	name := "Hello World"
	fmt.Printf("string is %s\n", name)
	printbytes(name)
	fmt.Printf("\n")
	printcharacter(name)
}
