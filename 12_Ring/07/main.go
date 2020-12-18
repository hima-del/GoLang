package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// Create a new ring of size 6
	r := ring.New(6)
	// Get the length of the ring
	n := r.Len()
	// Initialize the ring with some integer values
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}
	// Unlink three elements from r, starting from r.Next()
	r.Unlink(3)
	// Iterate through the remaining ring and print its contents
	r.Do(func(p interface{}) {
		fmt.Println(p.(int))
	})

}
