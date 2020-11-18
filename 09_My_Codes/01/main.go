package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("test.text")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	l, err := f.WriteString("hello world")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(l, "bytes written successfully")
}
