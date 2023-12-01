package main

import "fmt"

type course struct {
	name        string
	instructure string
	price       float64
}

func main() {
	
	c := course{name: "Basic Go", instructure: "AnuchitO", price: 2000}
	c2 := course{"Basic React", "AnuchitO", 2000}
	c3 := course{instructure: "A"}
	showCourse(c)
	showCourse(c2)
	showCourse(c3)
	c.price = 9999
	showCourse(c)
}

func showCourse(c course) {
	fmt.Printf("name: %s,name: %s,name: %.2f\n", c.name, c.instructure, c.price)
}
