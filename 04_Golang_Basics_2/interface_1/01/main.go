package main

import "fmt"

type vowelsFinder interface {
	findVowels() []rune
}

type myString string

func (ms myString) findVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

func main() {
	name := myString("John Doe")
	var v vowelsFinder
	v = name
	fmt.Printf("vowels are %c", v.findVowels())
}
