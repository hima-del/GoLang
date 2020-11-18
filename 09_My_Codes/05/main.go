//reading a file
package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("test.text")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("contents of file:", string(data))
}
