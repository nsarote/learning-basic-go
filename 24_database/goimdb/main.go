package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	_ "github.com/proullon/ramsql/driver"

	//"strconv"
	"math"
)

type DBType int32

const (
	MEM_DB DBType = iota
	POSTGRES_DB
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
	title := c.QueryParam("title")
	year := c.QueryParam("year")
	rating := c.QueryParam("rating")
	isSuperHero := c.QueryParam("isSuperHero")
	mList, err := getAllMovie(title, year, rating, isSuperHero)
	switch err {
	case nil:
		return c.JSON(http.StatusOK, mList)
	default:
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Server Error:" + err.Error()})
	}
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

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "goimdb"
)

func connPostgres() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	fmt.Println("Connect database success")
}

func connMemDB() {
	var err error
	db, err = sql.Open("ramsql", "goimdb")
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	fmt.Println("Connect database success")
}

func createTableMemDB() {
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
		fmt.Println("create table error:", err)
		return
	}

	fmt.Println("table created.")
}

func createTablePostgres() {
	fmt.Println("POSTGRES_DB")
	var err error
	createTb := `
	CREATE TABLE IF NOT EXISTS goimdb (
		id SERIAL PRIMARY KEY,
		imdbID TEXT NOT NULL UNIQUE,
		title TEXT NOT NULL,
		year INT NOT NULL,
		rating FLOAT NOT NULL,
		isSuperHero BOOLEAN NOT NULL
	);
	`
	_, err = db.Exec(createTb)
	if err != nil {
		fmt.Println("create table error:", err)
		return
	}

	fmt.Println("table created.")
}

func dropTablePostgres() {
	var err error
	dropTb := `DROP TABLE IF EXISTS goimdb;`
	_, err = db.Exec(dropTb)
	if err != nil {
		fmt.Println("drop table error:", err)
	}

	fmt.Println("drop table success.")
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

func round(val float32, precision int) float32 {
	return float32(math.Round(float64(val)*(math.Pow10(precision))) / math.Pow10(precision))
}

func insertData(imdbID, title string, year int, ratings float32, isSuperHero bool) (int64, error) {
	var insert string
	switch currentDBType {
	case MEM_DB:
		insert = `INSERT INTO goimdb(imdbid , title , year , rating, issuperhero) VALUES ($1 , $2 , $3 , $4 , $5 );`
	case POSTGRES_DB:
		insert = `INSERT INTO goimdb(imdbid , title , year , rating, issuperhero) VALUES ($1 , $2 , $3 , ROUND($4,1) , $5 );`
	}
	stmt, err := db.Prepare(insert)
	if err != nil {
		fmt.Println("prepare statement error:", err)
		log.Fatal("prepare statement error:", err)
		return 0, err
	}
	r, err := stmt.Exec(imdbID, title, year, round(ratings, 1), isSuperHero)
	if err != nil {
		fmt.Println("insert error:", err)
		log.Fatal("insert error:", err)
		return 0, err
	}
	l, err := r.LastInsertId()
	fmt.Println("lastINsertId", l, "err:", err)
	ef, err := r.RowsAffected()
	fmt.Printf("RowsAffecred %d err: %s\n", ef, err)

	return ef, nil
}

func getAllMovie(title, year, rating, isSuperHero string) ([]Movie, error) {
	var query = "SELECT id, imdbID, title, year, rating, isSuperHero FROM goimdb where 1=1 "
	if title != "" {
		//query += "and UPPER(title) like '%" + strings.ToUpper(title) + "%' "
		query += "and title = '" + title + "' "
	}
	if year != "" {
		query += "and year = " + year + " "
	}
	if rating != "" {
		//f, _ := strconv.ParseFloat(rating, 64)
		query += "and rating = " + rating + " "
	}
	if isSuperHero != "" {
		query += "and issuperhero = '" + isSuperHero + "' "
	}
	fmt.Println("query:", query)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("query error", err)
		return nil, err
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
			fmt.Println("for rows error", err)
			return nil, err
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
	return ms, nil
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
	rowx := db.QueryRow(`SELECT id, imdbID, title, year, rating, isSuperHero FROM goimdb WHERE imdbID=$1`, idForFind)
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

const currentDBType = MEM_DB

//const currentDBType = MEM_DB

func main() {
	defer db.Close()
	switch currentDBType {
	case MEM_DB:
		fmt.Println("MEM_DB")
		connMemDB()
		createTableMemDB()
	case POSTGRES_DB:
		fmt.Println("POSTGRES_DB")
		connPostgres()
		dropTablePostgres()
		createTablePostgres()
	default:
		fmt.Println("MEM_DB")
		connMemDB()
		createTableMemDB()
	}

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
