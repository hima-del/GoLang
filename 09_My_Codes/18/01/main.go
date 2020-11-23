//A never-ending football game.
package main

import (
	"fmt"
	"time"
)

func main() {
	var ball = make(chan string)
	kickBall := func(player string) {
		for {
			fmt.Println(<-ball, "kicked the ball")
			time.Sleep(time.Second)
			ball <- player
		}
	}
	go kickBall("john")
	go kickBall("alice")
	go kickBall("bob")
	go kickBall("emily")
	ball <- "referee"
	//var c chan bool
	//<-c
	//fmt.Println("channel", d)
	//time.Sleep(2 * time.Second)
}
