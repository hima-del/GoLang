package main

import (
	"fmt"
	"time"
)

func main() {
	m, _ := time.ParseDuration("1h30m")
	fmt.Printf("the movie is %v minutes long", m.Minutes())
}
