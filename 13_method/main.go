package main

import "fmt"

type course struct {
	name, instructure string
	price             int
}

// Receiver method - function include in type
func (c course) discount(d int) int {
	p := c.price - d
	fmt.Println("discount:", p)
	return p
}

func main() {
	c := course{name: "Basic Go", instructure: "AnuchitO", price: 2000}
	showCourse(c)
	d := c.discount(599)
	fmt.Println("discount price:", d)
	c.info()
}

func showCourse(c course) {
	fmt.Printf("name: %s,name: %s,name: %d\n", c.name, c.instructure, c.price)
}

func (c course) info() {
	fmt.Printf("name: %s,name: %s,name: %d\n", c.name, c.instructure, c.price)
}

/* func discount(c course) int {
	p := c.price - 599
	fmt.Println("discount:", p)
	return p
}*/
