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
	m1.info()
	m2.info()
}

func (m movie) info() {
	//fmt.Printf("title: %s,year: %d,ratings: %.2f,genres: %v,isSuperHero: %t\n", m.title, m.year, m.ratings, m.genres, m.isSuperHero)
	//fmt.Printf("%#v\n", m)
	fmt.Printf("%s (%d) - %.2f\n", m.title, m.year, m.ratings)
	fmt.Println("Genres:")
	for _, g := range m.genres {
		fmt.Printf("\t%s\n", g)
	}
	fmt.Println()
}
