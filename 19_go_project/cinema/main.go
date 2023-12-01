package main

import (
	"fmt"
	"github.com/nsarote/cinema/movie"
	"github.com/nsarote/cinema/ticket"
)

func init() {
	fmt.Println("init: main")
}

func main() {
	movieName := movie.FindMovieName("tt4154796")
	fmt.Println("movieName:", movieName)
	movie.ReviewMovie(movieName, 8.4)
	ticket.BuyTicket(movieName)
}
