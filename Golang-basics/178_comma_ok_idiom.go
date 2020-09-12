package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		c <- 42
		close(c)
	}()
	v, ok := <-c
	fmt.Println(v, ok)
	v1, ok1 := <-c
	fmt.Println(v1, ok1)
}
