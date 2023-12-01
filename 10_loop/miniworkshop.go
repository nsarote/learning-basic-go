package main

import "fmt"

func main() {
	genres := [3]string{"Action", "Adventure", "Fantasy"}
	fmt.Println("===before for loop: ")
	showArray(genres)
	for i := 0; i < len(genres); i++ {
		genres[i] = "Movie: " + genres[i]
	}
	fmt.Println("===after for loop: ")
	showArray(genres)
}

func showArray(arr [3]string) {
	for i, genre := range arr {
		fmt.Printf("genre[%d]: %s\n", i, genre)
	}
}
