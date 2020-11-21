//The two goroutines in this example talk to each other through an unbuffered channel.
package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go func(ch chan<- int, x int) {
		//time.Sleep(2 * time.Second)
		ch <- x * x
	}(c, 3)
	//done := make(chan struct{})
	go func(ch <-chan int) {
		n := <-ch
		fmt.Println(n)
		//time.Sleep(2 * time.Second)
		//done <- struct{}{}
	}(c)
	time.Sleep(2 * time.Second)
	//<-done
	fmt.Println("bye")
}
