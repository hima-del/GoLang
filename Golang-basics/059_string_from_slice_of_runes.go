package main

import "fmt"

func main() {
	runeslice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	str := string(runeslice)
	fmt.Printf(str)
}
