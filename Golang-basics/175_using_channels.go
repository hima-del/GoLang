package main

import "fmt"

func main() {
	c := make(chan int)
	go foo(c)
	bar(c)
}

func foo(c chan<- int) {
	c <- 42
}
func bar(c <-chan int) {
	fmt.Println(<-c)
}
