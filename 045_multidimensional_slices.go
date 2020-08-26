package main

import "fmt"

func printslice(a [][]string) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s", v2)
		}
		fmt.Printf("\n")
	}
}

func main() {
	b := [][]string{
		{"cat", "dog"},
		{"rat", "cow"},
		{"lion", "tiger"},
		{"pig", "fox"},
	}
	printslice(b)
}
