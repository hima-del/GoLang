package main

import "fmt"

func main() {
	number := 50
	switch {
	case number < 10:
		fmt.Println("number is less than 10")
	case number < 60:
		fmt.Println("number is less than 60")
	case number < 100:
		fmt.Println("number is less than 100")
	}
}
