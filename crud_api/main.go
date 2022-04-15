package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func main() {
	const Port string = "8080"
	r := mux.NewRouter()

	movies = append(
		movies,
		Movie{
			ID:    "1",
			Isbn:  "438277",
			Title: "Movie One",
			Director: &Director{
				Firstname: "John",
				Lastname:  "Doe",
			},
		},
	)

	movies = append(
		movies,
		Movie{
			ID:    "2",
			Isbn:  "43455",
			Title: "Movie Two",
			Director: &Director{
				Firstname: "Steve",
				Lastname:  "Smith",
			},
		},
	)

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server on port %s\n", Port)
	log.Fatal(http.ListenAndServe(":"+Port, r))
} // main

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
} // getMovies

func getMovie(w http.ResponseWriter, r *http.Request) {
} // getMovie

func createMovie(w http.ResponseWriter, r *http.Request) {
} // createMovie

func updateMovie(w http.ResponseWriter, r *http.Request) {
} // updateMovie

func deleteMovie(w http.ResponseWriter, r *http.Request) {
} // deleteMovie
