package main

import (
	"encoding/base64"
	"fmt"
)

const foobar = `foo bar`

func main() {
	encoding := base64.StdEncoding
	encodedfoobar := make([]byte, encoding.EncodedLen(len(foobar)))
	encoding.Encode(encodedfoobar, []byte(foobar))
	fmt.Printf("%s", encodedfoobar)
}
