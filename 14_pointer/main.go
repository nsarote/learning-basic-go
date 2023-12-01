package main

import "fmt"

type course struct {
	name, instructure string
	price             int
}

func discount(c course) int {
	p := c.price - 599
	fmt.Println("discount:", p)
	return p
}
func discountAddr(c *course) int {
	p := c.price - 599 // Go auto dereferencr
	fmt.Println("discount:", c.price)
	return p
}

func main() {
	var price int = 9999
	var addr *int = &price

	fmt.Println(price, addr)
	fmt.Printf("%T\n", addr)

	*addr = 9400 // write value to address
	fmt.Println(price, addr)

	v := *addr // read value from address
	fmt.Println(v)

	changePrice(price)
	fmt.Println("after", price, addr)

	changePriceFromAddress(addr)
	fmt.Println("after2", price, addr)

	fmt.Println()
	c := course{name: "Basic Go", instructure: "AnuchitO", price: 9999}
	d := discount(c)
	fmt.Println("discount price:", d)
	fmt.Println("price:", c.price)

	fmt.Println()
	e := discountAddr(&c)
	fmt.Println("discount price:", e)
	fmt.Println("price:", c.price)

	fmt.Println()
	c1 := &course{name: "Basic Go", instructure: "AnuchitO", price: 9999} // ประกาศ pointer of course
	c2 := new(course)                                                     // ประกาศ pointer of course
	fmt.Println(c1)
	fmt.Println(c2)

	fmt.Println()
	var c3 *int // zero of pointer is nil
	fmt.Println(c3)
}

func changePrice(price int) {
	price = price - 599
	fmt.Println("change", price, &price)
}

func changePriceFromAddress(p *int) {
	*p = *p - 599 //Dereference
	fmt.Println("changePriceFromAddress", *p, &p)
}
