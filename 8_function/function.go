package main

import "fmt"

func main() {
	a, b := add(42.1, 13.2)
	fmt.Println(a, b)
	fmt.Println(swap("hello", "world"))
	fmt.Println(split(9))

	fmt.Printf("%T, result:%f\n", add_, add_(1.2, 2.3))

	fmt.Println("\n")
	fmt.Println(compute(add_))
	fmt.Println(compute(minus_))

	fmt.Println("\nHigher Order Function")
	increase1, current1 := adder()
	fmt.Println(increase1, current1)
	fmt.Println(increase1(), current1())

	fmt.Println("\n")
	increase, current := adder()
	p := increase()
	fmt.Println(p)

	p = increase()
	fmt.Println(p)

	q := current()
	fmt.Println(q)
}

func adder() (func() int, func() int) {
	sum := 0
	return func() int {
			sum = sum + 1
			return sum
		}, func() int {
			return sum
		}
}

func increase() int {
	return 1
}

func current() int {
	return 2
}

func compute(fn func(float64, float64) float64) float64 {
	v := fn(42, 13)
	fmt.Printf("%T, result:%f\n", fn, v)
	return v
}

func add(x float64, y float64) (float64, string) {
	result := x + y
	fmt.Printf("%f + %f = %f\n", x, y, result)
	return result, "hello"
}

var add_ = func(x, y float64) float64 {
	return x + y
}

var minus_ = func(x, y float64) float64 {
	return x - y
}

func swap(x string, y string) (string, string) {
	return y, x
}

func split(sum int) (x int, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
