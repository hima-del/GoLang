//writing bytes to a file
package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("bytes")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	d2 := []byte{65, 83, 32, 89, 87, 90}
	n2, err := f.Write(d2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n2, "bytes written successfully")
}
