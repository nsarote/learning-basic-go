package main

import "fmt"

func main() {
	today := "Tuesday"
	switch today {
	case "Saturday", "Sunday":
		fmt.Println("today is Weekend")

	case "Monday":
		fallthrough
	case "Tuesday":
		fallthrough
	case "Wednesday":
		fallthrough
	case "Thurday":
		fallthrough
	case "Friday":
		fmt.Println("today is Weekday")
	default:
		fmt.Printf("Unknown")
	}

	/////
	num := 1
	switch {
	case num > 0:
		fmt.Println("num is positive")
	case num < 0:
		fmt.Println("num is negative")
	default:
		fmt.Println("num is zero")
	}

	/////
	ratings := 8.4
	switch {
	case ratings < 5:
		fmt.Println("Disappointed ☹")
	case ratings >= 5 && ratings < 7:
		fmt.Println("Normal 😧")
	case ratings >= 7 && ratings < 10:
		fmt.Println("Good 🥰")
	default:
		fmt.Println("😱😱😱")

	}

}
