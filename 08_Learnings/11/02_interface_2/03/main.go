//embedding interface
package main

import "fmt"

type leavecalculator interface {
	calculateleave() int
}

type salarycalculator interface {
	displaysalary()
}

type employeeOperations interface {
	salarycalculator
	leavecalculator
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
	fmt.Printf("%s %s has a salary %d\n", e.firstname, e.lastname, (e.basicpay + e.pf))
}

func (e employee) calculateleave() int {
	return e.totalleaves - e.leavestaken
}

func main() {
	e := employee{
		firstname:   "John",
		lastname:    "Doe",
		basicpay:    6000,
		pf:          100,
		totalleaves: 30,
		leavestaken: 3,
	}
	var empop employeeOperations
	empop = e
	empop.displaysalary()
	fmt.Println("leaves:", empop.calculateleave())
}
