package main

import "fmt"

type worker interface {
	work()
}

type person struct {
	name string
}

func (p person) work() {
	fmt.Println(p.name, "is working")
}

func describe(w worker) {
	fmt.Printf("interface type %T value %v\n", w, w)
}

func main() {
	p := person{
		name: "John Doe",
	}
	var w worker
	w = p
	describe(w)
	w.work()
}
