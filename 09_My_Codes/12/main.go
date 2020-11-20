//Sends to a buffered channel are blocked only when the buffer is full.
//Similarly receives from a buffered channel are blocked only when the buffer is empty.
package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 2)
	ch <- "hello"
	ch <- "world"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
