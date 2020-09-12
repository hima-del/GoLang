//adding sleep
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for i := 0; i < 5; i++ {
		fmt.Println("foo:", i)
		time.Sleep(time.Duration(3 * time.Microsecond))
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 5; i++ {
		fmt.Println("bar:", i)
		time.Sleep(time.Duration(20 * time.Microsecond))
	}
	wg.Done()
}
