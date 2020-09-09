package main

import (
	"fmt"

	"github.com/hima-del/GoLang/tree/exercise/ninjalevel12/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
}
