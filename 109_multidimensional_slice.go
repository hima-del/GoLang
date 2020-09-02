package main

import "fmt"

func printslice(a [][]int) {
	for _, v1 := range a {
		//fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Printf("%d\t", v2)
		}
		fmt.Printf("\n")
	}
}

func main() {
	b := [][]int{
		{1, 2, 10},
		{3, 4, 11},
		{5, 6, 12},
	}
	printslice(b)
}
