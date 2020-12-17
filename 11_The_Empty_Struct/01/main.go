package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var s string
	var c complex128
	var i int
	var str struct{}
	type S struct {
		A struct{}
		B struct{}
	}
	var st S
	fmt.Println(unsafe.Sizeof(s))
	fmt.Println(unsafe.Sizeof(c))
	fmt.Println(unsafe.Sizeof(i))
	fmt.Println(unsafe.Sizeof(str))
	fmt.Println(unsafe.Sizeof(st))
}

//As the empty struct consumes zero bytes, it follows that it needs no padding.
//Thus a struct comprised of empty structs also consumes no storage
