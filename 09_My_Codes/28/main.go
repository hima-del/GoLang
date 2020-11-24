package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	const input = "285317 1394676 18936417134"
	scanner := bufio.NewScanner(strings.NewReader(input))
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = bufio.ScanWords(data, atEOF)
		if err == nil && token != nil {
			_, err = strconv.ParseInt(string(token), 10, 32)
		}
		return
	}
	scanner.Split(split)
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}
	err := scanner.Err()
	if err != nil {
		fmt.Printf("invalid input:%s", err)
	}
}
