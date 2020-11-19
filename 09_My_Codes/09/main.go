//concurrency is handled in go using goroutines and channels
//channels are pipes using which goroutines communicate

//data:=<-a  read from the chnnel a
//a<-data  write to the channel a
//When a data is sent to a channel, the control is blocked in the send statement until some other Goroutine reads from that channel.
// Similarly when data is read from a channel, the read is blocked until some Goroutine writes data to that channel.
package main

import (
	"fmt"
	"time"
)

func hello(done chan bool) {
	fmt.Println("hello goroutine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("hello goroutine is awake and going to write to done")
	done <- true
}

func main() {
	done := make(chan bool)
	fmt.Println("main is going to call hello goroutine")
	go hello(done)
	<-done
	fmt.Println("main recieved data")
}
