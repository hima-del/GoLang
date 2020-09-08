package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	go send(even, odd, quit)

	recieve(even, odd, quit)

	fmt.Println("about to exit")
}

func send(even, odd chan<- int, quit chan<- bool) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(quit)
}

func recieve(even, odd <-chan int, quit <-chan bool) {
	for {
		select {
		case v := <-even:
			fmt.Println("the value recieved from even channel", v)
		case v := <-odd:
			fmt.Println("the value recieved from odd channel", v)
		case i, ok := <-quit:
			if !ok {
				fmt.Println("from comma ok bit", i)
			} else {
				fmt.Println("from comma ok bit", i)
			}
		}
	}
}
