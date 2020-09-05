package main

import (
	"fmt"
	"sort"
)

type Person struct {
	first string
	Age   int
}

type Byname []Person

func (a Byname) Len() int           { return len(a) }
func (a Byname) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Byname) Less(i, j int) bool { return a[i].first < a[j].first }
func main() {
	p1 := Person{"James", 30}
	p2 := Person{"Anna", 28}
	p3 := Person{"Doe", 35}
	p4 := Person{"Sara", 44}
	people := []Person{p1, p2, p3, p4}
	fmt.Println(people)
	sort.Sort(Byname(people))
	fmt.Println(people)
}
