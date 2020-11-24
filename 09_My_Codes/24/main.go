package main

import (
	"fmt"
	"time"
)

func main() {
	h, _ := time.ParseDuration("4h30m")
	fmt.Printf("%T", h.Hours())
	fmt.Printf("got %v hours of work left.", h.Hours())
}
