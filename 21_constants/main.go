package main

import "fmt"

const Pi = 3.14

type day int

func main() {
	const world = "XYZ" // cannot reassign value(final in java)
	fmt.Println("Hello", world)
	fmt.Println("Happy", Pi, "Day")

	const truth = true
	fmt.Println("Go rules?", truth)

	/*const (
		Sunday = 1
		Monday = 2
		Tuesday = 3
		Wednesday = 4
		Thursday = 5
		Friday = 6
		Saturday = 7
	)*/
	// use iota to same as above
	/*const (
		_ = iota
		Sunday
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)*/

	const (
		_          = iota
		Sunday day = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	fmt.Println("Sunday:", Sunday)
	fmt.Println("Monday:", Monday)
	fmt.Println("Tuesday:", Tuesday)
	fmt.Println("Wednesday:", Wednesday)
	fmt.Println("Thursday:", Thursday)
	fmt.Println("Friday:", Friday)
	fmt.Println("Saturday:", Saturday)
	fmt.Printf("%T Saturday:%v\n", Saturday, Saturday)

	///////////////////////////////////
	// Varidict function -> ... ->varag in java
	skills("js", "go", "python", "java")
}

func skills(xs ...string) {
	/*fmt.Println(xs[0])
	fmt.Println(xs[1])
	fmt.Println(xs[2])*/

	for _, s := range xs { // xs is slice
		fmt.Println("skill:", s)
	}
}
