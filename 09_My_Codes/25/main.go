package main

import (
	"fmt"
	"time"
)

func main() {
	u, _ := time.ParseDuration("1s")
	fmt.Printf("one sec is %d milliseconds\n", u.Milliseconds())
}
