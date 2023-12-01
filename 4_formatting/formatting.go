package main

import "fmt"

func main() {
	var r rune = '😀'
	fmt.Printf("type: %T -- r: %c\n", r, r)

	message2, age2, price2, check2 := "Test", 11, 123.45, true // Multiple declaration
	fmt.Printf("type: %T -- message2: %s\n", message2, message2)
	fmt.Printf("type: %T -- age2: %d\n", age2, age2)
	fmt.Printf("type: %T -- price2: %f\n", price2, price2)
	fmt.Printf("type: %T -- check2: %t\n", check2, check2)

	fmt.Printf("type: %T -- check2: %#v\n", check2, check2) // %#v print ได้ทุกค่า

}