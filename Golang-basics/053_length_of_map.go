package main

import "fmt"

func main() {
	employeesalary := make(map[string]int)
	employeesalary["a"] = 12000
	employeesalary["b"] = 56000
	employeesalary["c"] = 56908
	fmt.Printf("length of map is %d", len(employeesalary))
}
