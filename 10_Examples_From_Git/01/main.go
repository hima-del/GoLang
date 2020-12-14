package main

import "fmt"

func funcwithinterface(emptyinterface interface{}) {
	switch t := emptyinterface.(type) {
	case string:
		fmt.Println("type : string")
	case int:
		fmt.Println("type : int")
	case bool:
		fmt.Println("type : bool")
	default:
		fmt.Printf("type : %v", t)
	}
	fmt.Printf("data : %#v\n", emptyinterface)
}

func main() {
	var emptyinterface = [3]interface{}{}
	emptyinterface[0] = 23
	emptyinterface[1] = "foobar"
	emptyinterface[2] = false
	fmt.Printf("data : %#v\n", emptyinterface)
	for _, v := range emptyinterface {
		funcwithinterface(v)
	}
}
