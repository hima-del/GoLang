// program with Array sort
package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("integer reverse sort")
	num := []int{50, 90, 30, 10, 50}
	sort.Sort(sort.IntSlice(num))
	fmt.Println(num)

	fmt.Println("string reverse sort")
	text := []string{"elephant", "cat", "dog", "cat", "tiger", "lion"}
	sort.Sort(sort.Reverse(sort.StringSlice(text)))
	fmt.Println(text)
}
