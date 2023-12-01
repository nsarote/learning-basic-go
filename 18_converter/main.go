package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 256
	fmt.Printf("type:%T, val: %d\n", i, i)

	var f float64 = float64(i)
	fmt.Printf("type:%T, val: %f\n", f, f)

	var u uint8 = uint8(f) // be carefule down castig overflow case
	fmt.Printf("type:%T, val: %d\n", u, u)

	str := "72"
	var num int
	// num = int(str) // canot convert string to int should use strconv
	num, err := strconv.Atoi(str)
	fmt.Println(num, err)

	str = "72x"
	num, err = strconv.Atoi(str)
	fmt.Println(num, err)

}
