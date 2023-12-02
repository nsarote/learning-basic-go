package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Movie struct {
	ImdbID      string  `json:"imdbID"`
	Title       string  `json:"title"`
	Year        int     `json:"year"`
	Rating      float32 `json:"rating"`
	IsSuperHero bool    `json:"isSuperHero"`
}

var movies []Movie

func movieHandler(w http.ResponseWriter, req *http.Request) {
	method := req.Method
	switch method {
	case "GET":
		resp, err := json.Marshal(movies)
		if err != nil {
			fmt.Fprintf(w, "error marshall: %v", err)
			return
		}
		//fmt.Fprintf(w, "%s", resp)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write(resp)
		return
	case "POST":
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "error : %v", err)
			return
		}
		fmt.Println("Create movies :", string(body))
		m := Movie{}
		err = json.Unmarshal(body, &m)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "error unmarshall: %v", err)
			return
		}
		fmt.Println("Movie :", m)
		movies = append(movies, m)
		resp, err := json.Marshal(movies)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error marshall: %v", err)
			return
		}
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "%s", resp)
		return
	case "PUT":
		fmt.Fprintln(w, "Update movies")
		return
	case "DELETE":
		fmt.Fprintln(w, "Delete movies")
		return
	default:
		fmt.Fprintf(w, "%s not support.", method)
		return
	}

}
func main() {
	http.HandleFunc("/movies", movieHandler)

	log.Fatal(http.ListenAndServe("localhost:2565", nil))
}
