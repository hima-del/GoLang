package main

import "fmt"

func main() {
	a()
}

func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}
