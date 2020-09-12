package main

import "fmt"

func printbytes(s string) {
	fmt.Printf("bytes:")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

func printcharacters(s string) {
	fmt.Printf("characters:")
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
}

func main() {
	name := "Señor"
	fmt.Printf("string is %s\n", name)
	printbytes(name)
	fmt.Printf("\n")
	printcharacters(name)
}
