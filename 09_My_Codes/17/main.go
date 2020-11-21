package main

import "fmt"

func main() {
	c := make(chan int, 2) ///a buffered channel
	c <- 2
	c <- 3
	close(c)
	fmt.Println(len(c), cap(c))
	x, ok := <-c
	fmt.Println(x, ok)
	fmt.Println(len(c), cap(c))
	x, ok = <-c
	fmt.Println(x, ok)
	fmt.Println(len(c), cap(c))
	x, ok = <-c
	fmt.Println(x, ok)
	fmt.Println(len(c), cap(c))
	x, ok = <-c
	fmt.Println(x, ok)
	fmt.Println(len(c), cap(c))
}
