package main

import "fmt"

type promotion interface {
	discount() int
}

type infoer interface {
	getName() string
	info()
}
func displayCource(val infoer) {
	val.info()
}


type presenter interface {
	//info()
	//discount() int
	
	infoer //Embeded interface
	promotion
}
func summary(val presenter) {
	fmt.Println("discount price:", val.discount())
	val.info()
}



/* course */
type course struct {
	name string
	price int
	percentDiscount int
}
func (c course) discount() int {
	return c.price * c.percentDiscount/100
}
func (c course) getName() string {
	return c.name
}
func (c course) info() {
	fmt.Printf("info %T ; %#v\n - course name : %s\n", c, c, c.getName())
}

func main() {
	// var v interface{} //Empty Interface - Can keep every type; Go version >1.8 Added Generic type any
	var v any
	v = 36
	// show(v.(int)) // assert data type
	show(v)
	add2(v)
	v = "hello"
	show(v)
	add2(v) // panic: interface conversion: interface {} is string, not int

	fmt.Println("\n====[check type]===")
	checkType(123)
	checkType(123.456)
	checkType("hello")
	checkType(true)

	fmt.Println("\n====[declare type promotion]===")
	cA := course{name: "A", price: 1000, percentDiscount: 10}
	cB := course{name: "B", price: 2000, percentDiscount: 20}
	sale(cA)
	displayCource(cA)
	summary(cA)

	sale(cB)
	displayCource(cB)
	summary(cB)
}

func show(val any) {
	fmt.Printf("show %T ; %#v\n", val, val)
}

func add2(val any) {
	x, ok := val.(int)
	if ok {
		x = x + 2
		show(x)
	} else {
		fmt.Println(val, "is not int")
	}

	str, ok := val.(string)
	if ok {
		str += " is string"
		show(str)
	} else {
		fmt.Println("val is not int")
	}

}

func checkType(val any) {
	switch v := val.(type) {
	case int:
		fmt.Println(v, "is int")
	case string:
		fmt.Println(v, "is string")
	case float64:
		fmt.Println(v, "is float")
	default:
		fmt.Println(v, "is other")
	}
}

func sale(val promotion) {
	fmt.Printf("sale discount: %d\n", val.discount())
}
