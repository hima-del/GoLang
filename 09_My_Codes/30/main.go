package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	const input = "work hard dream big and never give up"
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}
	err := scanner.Err()
	if err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	fmt.Printf("%d\n", count)
}
