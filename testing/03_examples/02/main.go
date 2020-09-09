package main

import (
	"fmt"

	"./mymathtwo"
)

func main() {
	fmt.Println("2 + 3 =", mymathtwo.Sum(2, 3))
	fmt.Println("4 + 7 =", mymathtwo.Sum(4, 7))
	fmt.Println("5 + 9 =", mymathtwo.Sum(5, 9))
}
