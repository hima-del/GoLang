package main

import "fmt"

func main() {
	employeesalary := make(map[string]int)
	employeesalary["a"] = 12000
	employeesalary["b"] = 23000
	employeesalary["c"] = 45000
	for key, value := range employeesalary {
		fmt.Printf("employeesalary[%s]=%d\n", key, value)
	}
}
