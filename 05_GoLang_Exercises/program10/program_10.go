//Write a program to check if a character is a vowel or consonant
package main

import "fmt"

func main() {
	fmt.Println("enter the character")
	var character string
	fmt.Scanln(&character)
	if check(character) == true {
		fmt.Println("chracter is vowel")
	} else {
		fmt.Println("character is consonent")
	}

}

func check(charecter string) bool {
	vowel := [5]string{"a", "e", "i", "o", "u"}
	for i := 0; i < len(vowel); i++ {
		if charecter == vowel[i] {
			return true
		}
	}
	return false
}
