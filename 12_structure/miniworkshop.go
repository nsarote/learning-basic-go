package main

import "fmt"

type movie struct {
	title       string
	year        int
	ratings     float32
	genres      []string
	isSuperHero bool
}

func main() {

	m1 := movie{"Avengers: Endgame", 2019, 8.4, []string{"Action", "Drama"}, true}
	m2 := movie{"Infinity War", 2018, 8.4, []string{"Action", "Sci-Fi"}, true}
	showMovie(m1)
	showMovie(m2)

	mvs := []movie{m1, m2}
	for _, m := range mvs {
		showMovie(m)
	}
}

func showMovie(m movie) {
	fmt.Printf("title: %s,year: %d,ratings: %.2f,genres: %v,isSuperHero: %t\n", m.title, m.year, m.ratings, m.genres, m.isSuperHero)
	//fmt.Printf("%#v\n", m)
}
