package main

import "fmt"

func main() {
	c1 := make(chan int)
	c2 := make(chan<- int) //send
	c3 := make(<-chan int) //recieve

	fmt.Printf("%T\n", c1)
	fmt.Printf("%T\n", c2)
	fmt.Printf("%T\n", c3)
}
