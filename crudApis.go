package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// CRUD = Create, Read, Update, Delete

// | Operation | HTTP Method | Example Endpoint | Description     |
// | --------- | ----------- | ---------------- | --------------- |
// | Create    | POST        | `/users`         | Add new user    |
// | Read      | GET         | `/users`         | Get all users   |
// | Read      | GET         | `/users/{id}`    | Get single user |
// | Update    | PUT / PATCH | `/users/{id}`    | Update user     |
// | Delete    | DELETE      | `/users/{id}`    | Delete user     |

func crudApis() {
	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/create", createUser)
	http.HandleFunc("/update", updateUser)
	http.HandleFunc("/delete", deleteUser)

	http.ListenAndServe(":8080", nil)

	// Using gin
	// r := gin.Default()

	// r.POST("/users", createUser)
	// r.GET("/users", getUsers)
	// r.PUT("/users/:id", updateUser)
	// r.DELETE("/users/:id", deleteUser)

	// r.Run(":8080")
}

type UserCrudApis struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = []UserCrudApis{}
var nextID = 1

// Create User (POST)
// w http.ResponseWriter → Used to send response back to client
// r *http.Request → Contains all request data (body, headers, method, etc.)
func createUser(w http.ResponseWriter, r *http.Request) { // This is an HTTP handler function.
	var user UserCrudApis
	json.NewDecoder(r.Body).Decode(&user) // r.Body → contains raw JSON sent by client
	// json.NewDecoder(r.Body) → Creates a JSON decoder, .Decode(&user) → Converts JSON into Go struct
	// We pass pointer so the decoder can modify the original struct.
	user.ID = nextID
	nextID++
	users = append(users, user)

	w.Header().Set("Content-Type", "application/json") // Sets HTTP response header.
	// This tells the client: "I am sending JSON data back."

	json.NewEncoder(w).Encode(user) // Converts user struct → JSON; Writes it to response (w)
} // http.HandleFunc("/users", createUser)

// Get All Users (GET)
func getUsers(w http.ResponseWriter, r *http.Request) {
	if json.NewDecoder(r.Body) != nil {
		return
	}
	json.NewEncoder(w).Encode(users)
}

// Update User (PUT)
func updateUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id") // r.URL → The request URL, .Query() → Gets query parameters
	// .Get("id") → Retrieves value of id
	id, _ := strconv.Atoi(idStr) // strconv.Atoi() converts string → int

	for i, u := range users {
		if u.ID == id {
			json.NewDecoder(r.Body).Decode(&users[i]) // Decode New Data Into That User
			users[i].ID = id
			json.NewEncoder(w).Encode(users[i]) // This converts updated user → JSON and sends response.
			return
		}
	}
	http.NotFound(w, r)
} // http.HandleFunc("/users/update", updateUser)

// Delete User (DELETE)
func deleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			w.Write([]byte("Deleted"))
			return
		}
	}
	http.NotFound(w, r)
}
