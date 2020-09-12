package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		ch <- doSomething(5)
	}()
	fmt.Println(<-ch)
}

func doSomething(x int) int {
	return x * 5
}
