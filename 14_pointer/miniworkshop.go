package main

import "fmt"

type movie struct {
	title       string
	year        int
	ratings     float32
	votes       []float64
	genres      []string
	isSuperHero bool
}

func (m movie) info() {
	//fmt.Printf("title: %s,year: %d,ratings: %.2f,genres: %v,isSuperHero: %t\n", m.title, m.year, m.ratings, m.genres, m.isSuperHero)
	//fmt.Printf("%#v\n", m)
	fmt.Printf("%s (%d) - %.2f\n", m.title, m.year, m.ratings)
	fmt.Printf("Votes: %v\n", m.votes)
	fmt.Println("Genres:")
	for _, g := range m.genres {
		fmt.Printf("\t%s\n", g)
	}
	fmt.Println()
}

func (m *movie) addVote(v float64) {
	m.votes = append(m.votes, v)
}

func main() {
	m1 := &movie{"Avengers: Endgame", 2019, 8.4, []float64{7, 8, 9, 10, 6, 9, 9, 10, 8}, []string{"Action", "Drama"}, true}
	fmt.Println("Before : ")
	m1.info()
	m1.addVote(8)
	fmt.Println("After addVote : ")
	m1.info()
}
