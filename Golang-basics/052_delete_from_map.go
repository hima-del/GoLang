package main

import "fmt"

func main() {
	employeesalary := make(map[string]int)
	employeesalary["a"] = 23000
	employeesalary["b"] = 56000
	employeesalary["c"] = 67000
	delete(employeesalary, "a")
	fmt.Println("map after deleting is ", employeesalary)
}
