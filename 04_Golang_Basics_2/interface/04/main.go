//empty interface
package main

import "fmt"

func describe(i interface{}) {
	fmt.Printf("type = %T , value = %v\n", i, i)
}

func main() {
	s := "hello world"
	describe(s)
	i := 45
	describe(i)
	person := struct {
		name string
	}{
		name: "John Doe",
	}
	describe(person)
}
