package main

import (
	"fmt"
	"math"
)

type Number interface {
	int | float64
}

func min[T Number](a, b T) T {
	if a < b {
		return a
	} else {
		return b
	}
}

func minFloat(a, b float64) float64 {
	return math.Min(a, b)
}
func main() {
	var a, b = 1.3, 2.3
	fmt.Println("minFloat:", minFloat(a, b))

	fmt.Println("min:", minFloat(1, 2))
	fmt.Println("min:", minFloat(1.1, 2.2))
}
