package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//HTTP server is: Lightweight, Concurrent (handles multiple requests using goroutines),
// Production-ready, Part of the standard library (no external dependency required)

// An HTTP server: Listens on a port (like :8080), Accepts HTTP requests (GET, POST, etc.)
// Sends back HTTP responses

// http.ResponseWriter: used to send responce back to client
// *http.Request
// fmt.Println(r.Method)      // GET, POST
// fmt.Println(r.URL.Path)    // /about
// fmt.Println(r.Header)      // Headers
// fmt.Println(r.Body)        // Request body
func httpServer() {
	// Simplest http server
	http.HandleFunc("/", simplestHttp) // Registers a route
	http.ListenAndServe(":8080", nil)  // Starts the server

	// Handling Multiple Routes
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.ListenAndServe(":8080", nil)

}

// Simplest http server
func simplestHttp(w http.ResponseWriter, r *http.Request) { // Function that handles request
	fmt.Fprintln(w, "Hello, World!")
}

// Handling Multiple Routes
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome Home")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "About Page")
}

// Handling HTTP Methods
func handlingMethods(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		fmt.Fprintln(w, "GET Request")
		return
	}

	if r.Method == http.MethodPost {
		fmt.Fprintln(w, "POST Request")
		return
	}

	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}

// Reading Query Parameters
// http://localhost:8080/search?name=kartik
// name := r.URL.Query().Get("name")
// fmt.Fprintln(w, "Hello", name)

// Reading Form Data (POST)
// r.ParseForm()
// name := r.FormValue("name")

// Working with JSON
type UserHttpSever struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	var user UserHttpSever

	json.NewDecoder(r.Body).Decode(&user)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// Creating Custom Server
// server := &http.Server{
// 	Addr:    ":8080",
// 	Handler: nil,
// }
// server.ListenAndServe()

// Concurrency in Go HTTP Server: Go automatically creates a new goroutine per request.
// You do NOT need to manually use goroutines., That’s why Go servers scale very well.

// ServeMux (Server Multiplexer) is a request router in Go’s net/http package.
// Matches incoming request URLs (paths)
// Decides which handler function should run
// Routes the request to the correct function
// Incoming Request → Router (ServeMux) → Correct Handler Function
// Using ServeMux (Router)
func serveMux() {
	mux := http.NewServeMux() // Create a Router
	//  It creates a routing table (map of paths → handler functions) It waits for you to register routes

	mux.HandleFunc("/", home)
	mux.HandleFunc("/about", about)

	http.ListenAndServe(":8080", mux)
}
