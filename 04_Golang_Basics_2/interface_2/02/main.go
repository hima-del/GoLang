//a type can implement more than one interface
package main

import "fmt"

type salarycalculator interface {
	displaysalary()
}

type leavecalculator interface {
	calculateleave() int
}

type employee struct {
	firstname   string
	lastname    string
	basicpay    int
	pf          int
	totalleaves int
	leavestaken int
}

func (e employee) displaysalary() {
	fmt.Printf("%s %s has salary %d\n", e.firstname, e.lastname, (e.basicpay + e.pf))
}

func (e employee) calculateleave() int {
	return e.totalleaves - e.leavestaken
}

func main() {
	e := employee{
		firstname:   "John",
		lastname:    "Doe",
		basicpay:    5000,
		pf:          200,
		totalleaves: 30,
		leavestaken: 5,
	}
	var s salarycalculator
	s = e
	s.displaysalary()
	var l leavecalculator
	l = e
	fmt.Println("leaves", l.calculateleave())
}
