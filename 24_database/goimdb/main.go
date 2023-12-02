package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"strconv"

	"github.com/labstack/echo/v4"
	_ "github.com/proullon/ramsql/driver"
)

type Movie struct {
	ID          int64   `json:"id"`
	ImdbID      string  `json:"imdbID"`
	Title       string  `json:"title"`
	Year        int     `json:"year"`
	Rating      float32 `json:"rating"`
	IsSuperHero bool    `json:"isSuperHero"`
}

func getAllMovieHandler(c echo.Context) error {
	//title := c.QueryParam("title")
	y := c.QueryParam("year")
	//rating := c.QueryParam("rating")
	//isSuperHero := c.QueryParam("isSuperHero")
	if y == "" {
		return c.JSON(http.StatusOK, getAllMovie())
	}
	year, err := strconv.Atoi(y)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ms := []Movie{}
	for _, movie := range getAllMovie() {
		if movie.Year == int(year) {
			ms = append(ms, movie)
		}
	}
	return c.JSON(http.StatusOK, ms)

}

func getMovieByIdHandler(c echo.Context) error {
	id := c.Param("id")

	movie, err := getMovieByImdbID(id)
	switch err {
	case nil:
		return c.JSON(http.StatusOK, movie)
	case sql.ErrNoRows:
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Not Found ImdbID: " + id})
	default:
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Server Error:" + err.Error()})
	}
}

func createMovieHandler(c echo.Context) error {
	m := new(Movie)
	if err := c.Bind(m); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "bad request:" + err.Error()})
	} else if m.ImdbID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "bad request imdbID is null or empty"})
	}
	//fmt.Println("Movie :", m)
	record, err := insertMovie(*m)
	switch {
	case err == nil && record == 1:
		mo, _ := getMovieByImdbID(m.ImdbID)
		return c.JSON(http.StatusCreated, mo)
	case err.Error() == "UNIQUE constriint violation":
		return c.JSON(http.StatusConflict, map[string]string{"message": "movie already exist"})
	default:
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Server Error:" + err.Error()})
	}

}

func updateMovieHandler(c echo.Context) error {
	var idForUpdate string = c.Param("id")
	m := new(Movie)
	if err := c.Bind(m); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	} else if idForUpdate == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "bad request imdbID is null or empty"})
	}
	//fmt.Println("Movie :", m)
	//movies = append(movies, *m)
	_, err := getMovieByImdbID(idForUpdate)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "imdbID " + idForUpdate + " not found."})
	}
	fmt.Println("err1 :", err)
	err = updateMovie(idForUpdate, *m)
	fmt.Println("err2 :", err)
	switch {
	case err == nil:
		mo, _ := getMovieByImdbID(idForUpdate)
		return c.JSON(http.StatusOK, mo)
	case err != nil:
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Server Error:" + err.Error()})
	default:
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Server Error:"})
	}

}

func deleteMovieHandler(c echo.Context) error {
	idForDelete := c.Param("id")
	_, err := getMovieByImdbID(idForDelete)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "imdbID " + idForDelete + " not found."})
	}
	record, err := deleteMovieByImdbID(idForDelete)
	if record != 1 {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Server Error:" + err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Delete Success"})
}

// start connect database

var db *sql.DB

func conn() {
	var err error
	db, err = sql.Open("ramsql", "goimdb")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connect database success")
}

func createTable() {
	var err error
	createTb := `
	CREATE TABLE IF NOT EXISTS goimdb (
	id INT AUTO_INCREMENT,
	imdbID TEXT NOT NULL UNIQUE,
	title TEXT NOT NULL,
	year INT NOT NULL,
	rating FLOAT NOT NULL,
	isSuperHero BOOLEAN NOT NULL,
	PRIMARY KEY (id)
	);
	`

	_, err = db.Exec(createTb)
	if err != nil {
		log.Fatal("create table error:", err)
		return
	}

	fmt.Println("table created.")
}

func insertMovie(m Movie) (int64, error) {
	return insertData(m.ImdbID, m.Title, m.Year, m.Rating, m.IsSuperHero)
}

func updateMovie(idForUpdate string, m Movie) error {
	return updateData(idForUpdate, m.Title, m.Year, m.Rating, m.IsSuperHero)
}

func insertDatas() {
	totalRecord, _ := insertData("tt4154796", "Avengers: Endgame", 2019, 8.4, true)
	record, _ := insertData("tt4154756", "Avengers: Infinity War", 2018, 8.4, true)
	totalRecord += record
	record, _ = insertData("tt1211837", "Doctor Strange", 2016, 7.5, false)
	totalRecord += record
	record, _ = insertData("tt1825683", "Black Panther", 2018, 7.3, false)
	totalRecord += record

	fmt.Println("insertDatas totalRecord :", totalRecord)
}

func insertData(imdbID, title string, year int, ratings float32, isSuperHero bool) (int64, error) {
	insert := `
	INSERT INTO goimdb(imdbID,title,year,rating,isSuperHero)
	VALUES (?,?,?,?,?);
	`
	stmt, err := db.Prepare(insert)
	if err != nil {
		fmt.Println("prepare statement error:", err)
		return 0, err
	}
	r, err := stmt.Exec(imdbID, title, year, ratings, isSuperHero)
	if err != nil {
		fmt.Println("insert error:", err)
		return 0, err
	}
	l, err := r.LastInsertId()
	fmt.Println("lastINsertId", l, "err:", err)
	ef, err := r.RowsAffected()
	fmt.Printf("RowsAffecred %d err: %s\n", ef, err)
	return ef, nil
}

func getAllMovie() []Movie {
	rows, err := db.Query(`
	SELECT id, imdbID, title, year, rating, isSuperHero
	FROM goimdb
	`)
	if err != nil {
		log.Fatal("query error", err)
	}

	ms := []Movie{}
	for rows.Next() {
		var id int64
		var imdbID, title string
		var year int
		var rating float32
		var isSuperHero bool

		err := rows.Scan(&id, &imdbID, &title, &year, &rating, &isSuperHero)
		if err != nil {
			log.Fatal("for rows error", err)
		}
		//fmt.Println("row:", id, imdbID, title, year, rating, isSuperHero)
		movie := &Movie{
			ID:          id,
			ImdbID:      imdbID,
			Title:       title,
			Year:        year,
			Rating:      rating,
			IsSuperHero: isSuperHero,
		}
		ms = append(ms, *movie)
	}
	return ms
}

func updateData(imdbIDForUpdate, title string, year int, ratings float32, isSuperHero bool) error {
	stmt2, err := db.Prepare(`UPDATE goimdb SET title=$1,year=$2,rating=$3,issuperhero=$4 WHERE imdbID=$5`)
	if err != nil {
		fmt.Println("prepare statement error:", err)
		return err
	} else {
		fmt.Println("craete update prepare statement success")
	}
	r2, err := stmt2.Exec(title, year, ratings, isSuperHero, imdbIDForUpdate)
	if err != nil {
		fmt.Println("update error : ", err)
		return err
	} else {
		fmt.Println("Update record success : ", r2)
		return nil
	}
}

func deleteMovieByImdbID(imdbID string) (int64, error) {
	stmt2, err := db.Prepare(`DELETE from goimdb WHERE imdbID=$1`)
	if err != nil {
		log.Fatal("prepare statement error:", err)
	} else {
		fmt.Println("craete delete prepare statement success")
	}
	r2, err := stmt2.Exec(imdbID)
	if err != nil {
		fmt.Println("update error", err)
		return 0, err
	} else {
		fmt.Println("Update record success : ", r2)
		ef, err := r2.RowsAffected()
		return ef, err
	}
}

func getMovieByImdbID(idForFind string) (Movie, error) {
	rowx := db.QueryRow(`SELECT id, imdbID, title, year, rating, isSuperHero FROM goimdb WHERE imdbID=?`, idForFind)
	var id int64
	var imdbID, title string
	var year int
	var rating float32
	var isSuperHero bool
	err := rowx.Scan(&id, &imdbID, &title, &year, &rating, &isSuperHero)
	if err != nil {
		fmt.Println("scan one rowx error", err)
		return Movie{}, err
	}

	fmt.Println("one rowx:", id, imdbID, title, year, rating, isSuperHero)

	return Movie{
		ID:          id,
		ImdbID:      imdbID,
		Title:       title,
		Year:        year,
		Rating:      rating,
		IsSuperHero: isSuperHero,
	}, nil
}

// end connect database

func main() {
	conn()
	createTable()
	insertDatas()

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
