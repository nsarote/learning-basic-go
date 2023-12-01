package main

import (
	"fmt"
	"github.com/nsarote/cinema/movie"
	"github.com/nsarote/cinema/ticket"
)

func main() {
	fmt.Println("my project")

	movieName := movie.FindMovieName("tt4154796")
	fmt.Println("movieName:", movieName)
	movie.ReviewMovie(movieName, 8.4)
	ticket.BuyTicket(movieName)
}