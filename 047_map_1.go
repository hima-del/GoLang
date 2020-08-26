package main

import "fmt"

func main() {
	employeesalary := make(map[string]int)
	employeesalary["Ram"] = 12000
	employeesalary["Seetha"] = 39000
	employeesalary["Ravan"] = 56000
	fmt.Println("employeesalary map contents", employeesalary)
}
