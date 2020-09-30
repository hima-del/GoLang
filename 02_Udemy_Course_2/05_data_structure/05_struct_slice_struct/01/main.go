package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type person struct {
	Name string
	Age  int
}

type car struct {
	Model string
	Doors int
}

type items struct {
	Wisdom    []person
	Transport []car
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	p1 := person{
		Name: "John Doe",
		Age:  45,
	}

	p2 := person{
		Name: "Anna",
		Age:  35,
	}

	p3 := person{
		Name: "Wiiliam",
		Age:  40,
	}

	p4 := person{
		Name: "Sara",
		Age:  50,
	}

	p5 := person{
		Name: "Jenna",
		Age:  30,
	}

	c1 := car{
		Model: "Corolla",
		Doors: 4,
	}

	c2 := car{
		Model: "F150",
		Doors: 2,
	}
	cars := []car{c1, c2}
	persons := []person{p1, p2, p3, p4, p5}

	data := items{
		Wisdom:    persons,
		Transport: cars,
	}
	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
