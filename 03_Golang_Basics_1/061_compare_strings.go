package main

import "fmt"

func comparestrings(string1 string, string2 string) {
	if string1 == string2 {
		fmt.Printf("%s and %s are equal\n", string1, string2)
	} else {
		fmt.Printf("%s and %s are not equal\n", string1, string2)
	}
}

func main() {
	string1 := "go"
	string2 := "go"
	comparestrings(string1, string2)
	string3 := "apple"
	string4 := "orange"
	comparestrings(string3, string4)
}
