package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("length is %d and capacity is %d\n", len(a), cap(a))
	a = append(a, 7)
	fmt.Printf("length is %d and capacity is %d\n", len(a), cap(a))
	a = append(a, 8)
	fmt.Printf("length is %d and capacity is %d\n", len(a), cap(a))
	a = append(a, 8)
	fmt.Printf("length is %d and capacity is %d\n", len(a), cap(a))
	a = append(a, 10)
	fmt.Printf("length is %d and capacity is %d\n", len(a), cap(a))
	a = append(a, 11)
	fmt.Printf("length is %d and capacity is %d\n", len(a), cap(a))
	a = append(a, 12)
	fmt.Printf("length is %d and capacity is %d\n", len(a), cap(a))
	a = append(a, 13)
	fmt.Printf("length is %d and capacity is %d\n", len(a), cap(a))
}
