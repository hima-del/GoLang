package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var b bytes.Buffer //a buffer needs no initialization
	b.Write([]byte("hello"))
	fmt.Fprintf(&b, "world")
	b.WriteTo(os.Stdout)
}
