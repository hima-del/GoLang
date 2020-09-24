package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("cpus:", runtime.NumCPU())
	fmt.Println("goroutines:", runtime.NumGoroutine())
	counter := 0
	const gr = 100
	var wg sync.WaitGroup
	wg.Add(gr)

	for i := 0; i < gr; i++ {
		go func() {
			v := counter
			//time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("goroutines", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("goroutines", runtime.NumGoroutine())
	fmt.Println("count", counter)
}
