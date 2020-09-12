package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{5, 3, 8, 16, 11, 90, 55, 33, 1}
	sort.Ints(s)
	fmt.Println(s)
}
