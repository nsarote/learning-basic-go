package main

import "fmt"

var msg3 = "Hello Message 3"
// msg4: = "Hello Message 4" // syntax error: non-declaration statement outside function body
func main() {
	var msg1 string = "Hello Message 1"
	msg2:= "Hello Message 2"
	fmt.Println("Hello Gopher!!!")
	fmt.Println("msg1: ", msg1)
	fmt.Println("msg2: ", msg2)
	fmt.Println("msg3: ", msg3)
	//fmt.Println(msg4)

	var age int = 55
	fmt.Println("age: ", age)

	var price float64 = 1234.45
	fmt.Println("price: ", price)

	var marrige bool = true
	fmt.Println("marrige: ", marrige)

	message2, age2, price2, check2 := "Test", 11, 123.45, true // Multiple declaration
	fmt.Println("message2:", message2)
	fmt.Println("age2:", age2)
	fmt.Println("price2:", price2)
	fmt.Println("check2:", check2)

}