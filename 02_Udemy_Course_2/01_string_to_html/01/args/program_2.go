package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	// var s, sep string
	// sep = "hello"
	// for i := 1; i < len(os.Args); i++ {
	// 	s += sep + os.Args[i]
	// }
	// fmt.Println(s)
	prgname := os.Args[0]
	fmt.Printf("the program name is %s\n", prgname)
	names := os.Args[1:]
	fmt.Println(reflect.TypeOf(names))
	for _, name := range names {

		fmt.Printf("Hello, %s!\n", name)
	}
}

//$ go build read_args.go
//$ ./read_args Jan Peter Lucia
//The program name is ./read_args
//[]string
//Hello, Jan!
//Hello, Peter!
//Hello, Lucia!
