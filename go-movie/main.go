package main

import (
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	
	"github.com/gorilla/mux"
)

type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
	
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
	for index, movie := range movies {
		if movie.ID == params["id"] {
            movies = append(movies[:index], movies[index+1:]...)
            break
        }
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, movie := range movies {
		if movie.ID == params["id"] {
            json.NewEncoder(w).Encode(movie)
            return
        }
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
    var movie Movie
    _ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, movie := range movies {
		if movie.ID == params["id"] {
            movies = append(movies[:index], movies[index+1:]...)
            var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
        }
	}
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{
		ID: "1", 
		Isbn: "1234567890", 
		Title: "Gone with the Wind", 
		Director: &Director{
			FirstName: "John", 
			LastName: "Steinbeck",
		},
	})
	movies = append(movies, Movie{
		ID: "2", 
		Isbn: "59815629", 
		Title: "Movie that is not a Movie", 
		Director: &Director{
			FirstName: "Jake", 
			LastName: "Skake",
		},
	})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie)
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	fmt.Print("Starting server...[8080]")
	log.Fatal(http.ListenAndServe(":8080", r))
}
