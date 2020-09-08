package main

import (
	"os"
)

func main() {
	//defer foo()
	_, err := os.Open("no-file.txt")
	if err != nil {
		//fmt.Println("err happened", err)
		//log.Println("err happened", err)
		//log.Fatalln(err) //equivalent to println() followed by a call to Exit(1)
		//log.Panicln(err)
		panic(err)
	}
}

//func foo() {
//	fmt.Println("when os.Exit() is called ,deferred functions don't run")
//}
