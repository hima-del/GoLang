package main

import "fmt"

type salarycalculator interface {
	calculatesalary() int
}

type permanent struct {
	empid    int
	basicpay int
	pf       int
}

type contract struct {
	empid    int
	basicpay int
}

func (p permanent) calculatesalary() int {
	return p.basicpay + p.pf
}

func (c contract) calculatesalary() int {
	return c.basicpay
}

func totalexpense(s []salarycalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.calculatesalary()
	}
	fmt.Printf("total expense is %d", expense)
}

func main() {
	pemp1 := permanent{
		empid:    1,
		basicpay: 5000,
		pf:       20,
	}
	pemp2 := permanent{
		empid:    2,
		basicpay: 6000,
		pf:       30,
	}
	cemp1 := contract{
		empid:    3,
		basicpay: 3000,
	}
	employees := []salarycalculator{pemp1, pemp2, cemp1}
	totalexpense(employees)
}
