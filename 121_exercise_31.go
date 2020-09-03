package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourwheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t := truck{
		fourwheel: true,
		vehicle: vehicle{
			doors: 2,
			color: "yellow",
		},
	}
	s := sedan{
		luxury: true,
		vehicle: vehicle{
			doors: 4,
			color: "red",
		},
	}
	fmt.Println(t)
	fmt.Println(s)
	fmt.Println(t.fourwheel)
	fmt.Println(t.doors)
	fmt.Println(t.color)
	fmt.Println(s.luxury)
	fmt.Println(s.doors)
	fmt.Println(s.color)
}
