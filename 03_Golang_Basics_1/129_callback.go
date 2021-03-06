package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6}
	y := sum(xi...)
	fmt.Println(y)
	s2 := even(sum, xi...)
	fmt.Println("even numbers", s2)
	s3 := odd(sum, xi...)
	fmt.Println("odd numbers", s3)
}

func sum(x ...int) (total int) {
	total = 0
	fmt.Printf("%T\n", x)
	for _, v := range x {
		total += v
	}
	return
}

func even(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

func odd(f func(xi ...int) int, ui ...int) int {
	var yi []int
	for _, v := range ui {
		if v%2 != 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

//142_exercise41
