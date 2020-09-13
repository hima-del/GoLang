//program to check wheather a channel is closed or not
package main

import "fmt"

func main() {
	c := make(chan int)
	fmt.Println(isClosed(c))
	close(c)
	fmt.Println(isClosed(c))
}

func isClosed(ch <-chan int) bool {
	select {
	case <-ch:
		return true
	default:
	}
	return false
}
