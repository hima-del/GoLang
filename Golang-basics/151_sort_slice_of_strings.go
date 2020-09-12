package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"mango", "banana", "pine apple", "apple", "orange", "grapes"}
	sort.Strings(s)
	fmt.Println(s)
}
