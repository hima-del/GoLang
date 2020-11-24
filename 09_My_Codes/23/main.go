//func shuffle
package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	words := strings.Fields("fly like eagles")
	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})
	fmt.Println(words)
}
