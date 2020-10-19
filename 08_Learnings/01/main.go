//in this code i learn about Query() in the URL package
package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	u, err := url.Parse("https://example.org/?a=1&a=2&b=&=3")
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	fmt.Println(q["a"])
	fmt.Println(q.Get("a"))
	fmt.Println(q.Get("b"))
	fmt.Println(q.Get(""))
}
