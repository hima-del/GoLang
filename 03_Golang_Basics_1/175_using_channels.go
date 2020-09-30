package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go foo(c)
	go bar(c)
	time.Sleep(time.Second) //wait for second otherwise the main go routine will end before the completion of foo and bar
}

func foo(c chan<- int) {
	c <- 42
}
func bar(c <-chan int) {
	fmt.Println(<-c)
}
