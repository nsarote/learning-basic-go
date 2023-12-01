package movie

import "fmt"

func init() {
	fmt.Println("init: movie")
}

func ReviewMovie(name string, rating float64) {
	fmt.Printf("I reviewed %s and it's rating %f\n", name, rating)
}

func FindMovieName(imdbID string) string {
	switch imdbID {
	case "tt4154796":
		return "Avengers: Endgame"
	case "tt1825683":
		return "Black Panther"
	default:
		return "not found"
	}
}