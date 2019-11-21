package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"strconv"
)

// Song represents a song title and artist associated with a unique ID
type Song struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
}

var songs []Song

func getSongs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(songs)
}

func createSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var song Song
	_ = json.NewDecoder(r.Body).Decode(&song)
	song.ID = strconv.Itoa(rand.Intn(1000000))
	songs = append(songs, song)
	json.NewEncoder(w).Encode(&song)
}

func updateSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range songs {
		songs = append(songs[:index], songs[index+1:]...)

		if item.ID == params["id"] {
			var song Song
			_ = json.NewDecoder(r.Body).Decode(&song)
			song.ID = params["id"]
			songs = append(songs, song)
			json.NewEncoder(w).Encode(&song)

			break
		}
	}
	json.NewEncoder(w).Encode(songs)
}

func getSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range songs {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			break
		}
	}
}

func deleteSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range songs {
		if item.ID == params["id"] {
			songs = append(songs[:index], songs[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(&songs)
}

func testyMctesterson() string {
	return "YOU DIDN'T TEST ME"
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/songs", getSongs).Methods("GET")
	router.HandleFunc("/songs", createSong).Methods("POST")
	router.HandleFunc("/songs/{id}", updateSong).Methods("PUT")
	router.HandleFunc("/songs/{id}", getSong).Methods("GET")
	router.HandleFunc("/songs/{id}", deleteSong).Methods("DELETE")

	http.ListenAndServe(":8000", router)
}
