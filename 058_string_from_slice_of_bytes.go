package main

import "fmt"

func main() {
	//byteslice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	byteSlice := []byte{67, 97, 102, 195, 169}
	str := string(byteSlice)
	fmt.Printf(str)
}
