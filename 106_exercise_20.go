package main

import "fmt"

func main() {
	var favsport string = "cricket"
	switch favsport {
	case "football":
		fmt.Println("favourite sport is football")
	case "hockey":
		fmt.Println("favourite sport is hockey")
	case "cricket":
		fmt.Println("favourite sport is cricket")
	}

}
