package main

import (
	"fmt"
	"errors"
)

// type error interface { // Defind in go already
// 	Error() string
// }
var myErr = errors.New("My custom error message")

type MyError struct{}

func (e MyError) Error() string {
	return "MyError"
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		// err := MyError{}
		// err := fmt.Errorf("Divide by zero")
		// err := fmt.Errorf("Can't devide %f by zero", a)
		err := myErr
		return -0.0, err
	}
	return a / b, nil
}
func main() {
	r, error := Divide(21, 7)
	fmt.Printf("r: %.2f; %v\n", r, error)

	r, error = Divide(1, 0)
	fmt.Printf("r: %.2f; %v\n", r, error)
	if error != nil {
		fmt.Println("handle error :", error)
		return
	}
}
