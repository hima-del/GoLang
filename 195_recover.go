package main

import "fmt"

func main() {
	f()
	fmt.Println("returned normally from f")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered in f", r)
		}
	}()
	fmt.Println("calling g")
	g(0)
	fmt.Println("returned normally from g")
}

func g(i int) {
	if i > 3 {
		fmt.Println("panicking")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("printing in g", i)
	g(i + 1)
}
