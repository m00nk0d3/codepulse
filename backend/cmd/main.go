package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/m00nk0d3/codePulse/internal/handlers"
)

func init() {
	// load .env

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	r := mux.NewRouter()

	// Definição das rotas
	r.HandleFunc("/repos", getRepos).Methods("GET")
	r.HandleFunc("/repos/{repo}/commits", getCommits).Methods("GET")
	r.HandleFunc("/repos/{repo}/pulls", getPullRequests).Methods("GET")
	r.HandleFunc("/notifications", getNotifications).Methods("GET")
	r.HandleFunc("/login", handlers.HandleGitHubLogin).Methods("GET")
	r.HandleFunc("/callback", handlers.HandleGitHubCallback).Methods("GET")
	// Inicia o servidor

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
