//zero value of interface
package main

import "fmt"

type describer interface {
	describe()
}

func main() {
	var d1 describer
	if d1 == nil {
		fmt.Printf("d1 is nil and has a type %T value %v", d1, d1)
	}
}
