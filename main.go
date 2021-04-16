package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

//  create struct same like ES6 / java class

type Post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

// initialise empty list
var posts []Post

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/posts", getPosts).Methods("GET")

	http.ListenAndServe(":8000", router)

}
