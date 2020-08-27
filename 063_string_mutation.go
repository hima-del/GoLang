package main

import "fmt"

func mutate(s []rune) string {
	s[0] = 'a'
	return string(s)
}

func main() {
	name := "hello"
	fmt.Printf(mutate([]rune(name)))
}
