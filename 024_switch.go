package main

import "fmt"

func main() {
	var finger int = 4
	fmt.Printf("Finger %d is ", finger)
	switch finger {
	case 1:
		fmt.Println("thumb")
	case 2:
		fmt.Println("index")
	case 3:
		fmt.Println("middle")
	case 4:
		fmt.Println("ring")
	case 5:
		fmt.Println("small")
	}
}
