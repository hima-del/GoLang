package main

import "fmt"

func main() {
	a := [5]float64{55.6, 78, 99.3, 33, 44}
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d th element is %.1f\n", i, a[i])
	}
}
