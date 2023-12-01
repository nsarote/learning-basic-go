package main

import (
	"fmt"
	"math"
)

func main() {
	num := 34
	if num == 34 {
		fmt.Println("Yes, it's Thirty four")
	} else {
		fmt.Println("No, it's Thirty four")
	}

	if num == 34 {
		fmt.Println("Yes, it's Thirty four")
	} else {
		fmt.Println("No, it's Thirty four")
	}

	limit := 225.0
	p := 3
	v := math.Pow10(p)
	fmt.Println("v: ", v)
	if v >= limit {
		fmt.Printf("10 power %d is %.0f over limit %.0f\n", p, v, limit)
	} else {
		fmt.Printf("10 power %d is ", p, v)
	}


	ratings := 8.4
	if ratings < 5 {
		fmt.Println("Disappointed ☹")
	} else if ratings >= 5 && ratings < 7 {
		fmt.Println("Normal 😧")
	} else if ratings >= 7 && ratings < 10 {
		fmt.Println("Good 🥰")
	} else {
		fmt.Println("😱😱😱")
	}
}
