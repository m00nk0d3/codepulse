package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Definição das rotas
	r.HandleFunc("/repos", getRepos).Methods("GET")
	r.HandleFunc("/repos/{repo}/commits", getCommits).Methods("GET")
	r.HandleFunc("/repos/{repo}/pulls", getPullRequests).Methods("GET")
	r.HandleFunc("/notifications", getNotifications).Methods("GET")

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func getRepos(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List of repositories"))
}

func getCommits(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Commit history"))
}

func getPullRequests(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pull requests with diffs"))
}

func getNotifications(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GitHub notifications"))
}
