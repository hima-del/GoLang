package main

import "fmt"

func main() {
	var number int = 15
	switch {
	case number > 5 && number <= 10:
		fmt.Printf("%d is greater than 5 and less than or equal to 10", number)
	case number > 11 && number <= 20:
		fmt.Printf("%d is greater than 11 and less than or equal to 20", number)
	case number >= 21:
		fmt.Printf("%d id gt=reater than or equal to 21", number)
	}
}
