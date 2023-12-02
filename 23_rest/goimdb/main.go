package main

import (
	//"io/ioutil"
	"fmt"
	"log"
	"net/http"

	"strconv"

	"github.com/labstack/echo/v4"
)

type Movie struct {
	ImdbID      string  `json:"imdbID"`
	Title       string  `json:"title"`
	Year        int     `json:"year"`
	Rating      float32 `json:"rating"`
	IsSuperHero bool    `json:"isSuperHero"`
}

var movies = []Movie{
	{
		ImdbID:      "tt4154796",
		Title:       "Avengers: Endgame",
		Year:        2019,
		Rating:      8.4,
		IsSuperHero: true,
	},
	{
		ImdbID:      "tt4154795",
		Title:       "HULK",
		Year:        2018,
		Rating:      8.3,
		IsSuperHero: true,
	},
}

func getAllMovieHandler(c echo.Context) error {
	y := c.QueryParam("year")
	if y == "" {
		return c.JSON(http.StatusOK, movies)
	}
	year, err := strconv.Atoi(y)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ms := []Movie{}
	for _, movie := range movies {
		if movie.Year == int(year) {
			ms = append(ms, movie)
		}
	}
	return c.JSON(http.StatusOK, ms)

}

func getMovieByIdHandler(c echo.Context) error {
	id := c.Param("id")

	for _, movie := range movies {
		if movie.ImdbID == id {
			return c.JSON(http.StatusOK, movie)
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"message:": "Not Found ImdbID: " + id})
}

func createMovieHandler(c echo.Context) error {
	m := new(Movie)
	if err := c.Bind(m); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	fmt.Println("Movie :", m)
	movies = append(movies, *m)
	return c.JSON(http.StatusCreated, m)
}

func updateMovieHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "")
}

func deleteMovieHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "")
}

func main() {
	e := echo.New()

	e.GET("/movies", getAllMovieHandler)
	e.GET("/movies/:id", getMovieByIdHandler)
	e.POST("/movies", createMovieHandler)
	e.PUT("/movies/:id", updateMovieHandler)
	e.DELETE("/movies/:id", deleteMovieHandler)

	port := "2565"
	log.Println("starting... port:", port)
	log.Fatal(e.Start(":" + port))
}
