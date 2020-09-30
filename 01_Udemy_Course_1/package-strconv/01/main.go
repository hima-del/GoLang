package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := "10"
	s, err := strconv.Atoi(i)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T, %v", s, s)
}
