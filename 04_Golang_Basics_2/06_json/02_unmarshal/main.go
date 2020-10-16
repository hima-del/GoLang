package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"First":"John","Last":"Doe","Age":44},{"First":"anna","Last":"jackson","Age":34}]`
	bs := []byte(s)
	people := []person{}
	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(people)

	for i, v := range people {
		fmt.Println("person number ", i)
		fmt.Println(v.First, v.Last, v.Age)
	}
}
