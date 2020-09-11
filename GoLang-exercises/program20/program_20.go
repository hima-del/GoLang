//perefct number or not
package main

import "fmt"

func main() {
	fmt.Println("enter the number")
	var number int
	fmt.Scanln(&number)
	perfectcheck(number)
}

func perfectcheck(n int) {
	var a []int
	var sum int = 0
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			a = append(a, i)
		}
	}
	//fmt.Printf("%T", a)
	for _, v := range a {
		sum += v
	}
	if n == sum {
		fmt.Println("number is perfect number")
	} else {
		fmt.Println("number is not perfect number")
	}
}
