package main

import "fmt"

func charsandbyteposition(s string) {
	for index, rune := range s {
		fmt.Printf("%c character is at byte %d\n", rune, index)
	}
}

func main() {
	name := "Señor"
	charsandbyteposition(name)
}
