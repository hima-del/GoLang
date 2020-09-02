package main

import "fmt"

func main() {
	a := []string{"James", "Bond", "Shaken, not stirred"}
	b := []string{"Miss", "Moneypenny", "Helloooo, James"}
	c := [][]string{a, b}
	fmt.Println(c)
	for _, v1 := range c {
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
		fmt.Printf("\n")
	}
}
