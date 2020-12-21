package main

import "fmt"

type describer interface {
	describe()
}

type person struct {
	name string
	age  int
}

func (p person) describe() {
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

type address struct {
	state   string
	country string
}

func (a *address) describe() {
	fmt.Printf("state is %s country is %s", a.state, a.country)
}

func main() {
	var d1 describer
	p1 := person{"John Doe", 25}
	d1 = p1
	d1.describe()
	p2 := person{"James Jo", 34}
	d1 = &p2
	d1.describe()

	var d2 describer
	a := address{"Auckland", "NZ"}
	d2 = &a
	d2.describe()
}
