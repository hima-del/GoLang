package main

import "fmt"

func main() {
	employeesalary := make(map[string]int)
	employeesalary["a"] = 12000
	employeesalary["b"] = 90040
	employeesalary["c"] = 45000

	employee := "a"
	salary := employeesalary["a"]
	fmt.Printf("%s has a salary %d", employee, salary)
}
