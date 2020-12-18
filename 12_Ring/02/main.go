package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// Create a new ring of size 4
	r := ring.New(4)
	fmt.Println(r.Len())
}
