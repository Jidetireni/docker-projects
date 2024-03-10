package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Article represents a single article with Title, Description, and Content.
type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// Articles represents a collection of Article objects.
type Articles []Article

// allArticles returns all articles as JSON response.
func allArticles(w http.ResponseWriter, r *http.Request) {
	// Create a slice of Article objects with one example article.
	articles := Articles{
		Article{Title: "Test Title", Desc: "Test Description", Content: "Hello world"},
	}

	// Print a message to console indicating the endpoint was hit.
	fmt.Println("Endpoint hit: All article Endpoint")

	// Encode the articles slice as JSON and write it as the HTTP response.
	json.NewEncoder(w).Encode(articles)
}

// homePage responds with a simple message indicating the homepage was hit.
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

// handleRequests sets up HTTP routes and starts the server.
func handleRequests() {
	// Define HTTP routes and their corresponding handler functions.
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)

	// Start the HTTP server on port 8081.
	log.Fatal(http.ListenAndServe(":8081", nil))
}

// The entry point of the application.
func main() {
	// Call handleRequests to set up routes and start the server.
	handleRequests()
}

