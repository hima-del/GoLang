//length vs capacity
//The capacity of a buffered channel is the number of values that the channel can hold
//The length of the buffered channel is the number of elements currently queued in it.
package main

import "fmt"

func main() {
	ch := make(chan string, 3)
	ch <- "hello"
	ch <- "world"
	fmt.Println("capacity is ", cap(ch))
	fmt.Println("length is ", len(ch))
	fmt.Println("read value ", <-ch)
	fmt.Println("new length is ", len(ch))
	fmt.Println("read value ", <-ch)
	fmt.Println("new length is ", len(ch))
}
