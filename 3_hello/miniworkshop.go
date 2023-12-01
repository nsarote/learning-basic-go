package main

import "fmt"

func main() {
	title, year, ratings, genre, isSuperHero := "Avengers: Endgame", 2019, 8.4, "Sci-Fi", true
	fmt.Println("เรื่อง: ", title)
	fmt.Println("ปี: ", year)
	fmt.Println("เรตติ้ง: ", ratings)
	fmt.Println("ประเภท: ", genre)
	fmt.Println("ซุปเปอร์ฮีโร่: ", isSuperHero)

	movie := `เรื่อง:  Avengers: Endgame
ปี:  2019
เรตติ้ง:  8.4
ประเภท:  Sci-Fi
ซุปเปอร์ฮีโร่:  true`
	fmt.Println("\nmovie: ", movie)

	fav := '⭐'
	fmt.Printf("\nเรื่อง: %s\n", title)
	fmt.Printf("ปี: %d\n", year)
	fmt.Printf("เรตติ้ง: %.1f\n", ratings)
	fmt.Printf("ประเภท: %s\n", genre)
	fmt.Printf("ซุปเปอร์ฮีโร่: %t\n", isSuperHero)
	fmt.Printf("Favorite: %c\n", fav)
}
