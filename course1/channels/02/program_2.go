//program to check wheather a channel is closed or not
package main

import "fmt"

func main() {
	c := make(chan int)
	fmt.Println(isClosed(c))
	close(c) //when the channel is closed you cannnot pou values onto the channel but you can still recieve values from that channel
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
