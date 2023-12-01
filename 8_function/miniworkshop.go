package main

import "fmt"

func main() {
	fmt.Println(emote(4.9))
	fmt.Println(emote(5.0))
	fmt.Println(emote(7.0))
	fmt.Println(emote(15.0))
}

func emote(ratings float64) string {
	switch {
	case ratings < 5:
		return "Disappointed 🙂"
	case ratings >= 5 && ratings < 7:
		return "Normal 😧"
	case ratings >= 7 && ratings < 10:
		return "Good 🥰"
	default:
		return "😱😱😱"
	}
}
