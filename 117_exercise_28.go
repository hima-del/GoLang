package main

import "fmt"

func main() {
	detail := map[string][]string{
		"bond_james":      []string{"shaken, not stirred", "martinis", "women"},
		"moneypenny_miss": []string{"james bond", "literature", "computer science"},
		"no_dr":           []string{"being evil", "ice cream", "sunsets"},
	}
	fmt.Println(detail)
	detail["james_annie"] = []string{"agar agar", "chocolate", "football"}
	delete(detail, "no_dr")
	for key, value1 := range detail {
		fmt.Println("this is the record of ", key)
		for i, value2 := range value1 {
			fmt.Println(i, value2)
		}
	}
}
