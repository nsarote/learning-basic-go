package main

import "fmt"

func main() {
	xs := []float64{4, 5, 7, 8, 3, 8, 0}
	ys := []float64{7, 2, 10, 9, 7}
	fmt.Println("xs")
	show(xs)
	fmt.Println("ys")
	show(ys)

	votes := append(xs, ys...)
	fmt.Println("votes")
	show(votes)

	fmt.Println("votes[5:9]")
	show(votes[5:9])
}

func show(slice []float64) {
	fmt.Printf("length: %d => %#v\n", len(slice), slice)
}
