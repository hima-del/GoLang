package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(strings.NewReader("gopher"))
	for scanner.Scan() {
		fmt.Println(len(scanner.Bytes()) == 6)
	}
	err := scanner.Err()
	if err != nil {
		fmt.Fprintln(os.Stderr, "shoulsn't see an error scanning a string")
	}
}
