package main

import "fmt"

func main() {
	c := make(chan int, 2)
	c <- 42 //send to channel
	c <- 43
	fmt.Println(<-c) //recieve from channel
	fmt.Println(<-c)
}
