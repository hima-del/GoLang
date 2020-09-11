//enter a string and find if it is pallindrome or not
package main

import "fmt"

func main() {
	fmt.Println("enter a string")
	var str string
	fmt.Scanln(&str)
	answer := reversestring(str)
	fmt.Println("reversed string is", answer)
	if isPallindrome(str) == true {
		fmt.Println("string is pallindrome")
	} else {
		fmt.Println("string is not pallindrome")
	}
}

func reversestring(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func isPallindrome(s string) bool {
	if s == reversestring(s) {
		return true
	}
	return false
}
