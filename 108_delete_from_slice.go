package main

import "fmt"

func main() {
	x := []string{"cat", "dog", "rat", "cow", "goat", "bull", "tiger", "rabbit"}
	fmt.Println(x)
	x = append(x, "elephant", "lion", "deer", "beer", "fox")
	fmt.Println(x)
	y := append(x[:3], x[5:]...)
	fmt.Println(y)
}
