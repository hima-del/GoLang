package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

type person struct {
	Name string
	Age  int
}

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

type items struct {
	Wisdom    []person
	Transport []car
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 3 {
		s = s[:3]
	}
	return s
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
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}

}
