package main

import "fmt"

func main() {
	employeesalary := make(map[string]int)
	employeesalary["a"] = 12000
	employeesalary["b"] = 34000
	employeesalary["c"] = 45000
	newemployee := "d"
	value, ok := employeesalary[newemployee]
	if ok == true {
		fmt.Printf("employ %s is found and salary is %d", newemployee, value)
	}
	fmt.Printf("employee %s is not found and salary is %d", newemployee, value)
}
