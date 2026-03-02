package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

// replication of bank transactions using crud apis
func main() {

	http.HandleFunc("/accounts", createAccount)

	http.HandleFunc("/accounts/", func(w http.ResponseWriter, r *http.Request) {

		if strings.HasSuffix(r.URL.Path, "/deposit") {
			deposit(w, r)
			return
		}

		if strings.HasSuffix(r.URL.Path, "/withdraw") {
			withdraw(w, r)
			return
		}

		if strings.HasSuffix(r.URL.Path, "/transactions") {
			getTransactions(w, r)
			return
		}

		if r.Method == "GET" {
			getAccount(w, r)
			return
		}
	})

	http.ListenAndServe(":8080", nil)
}

// Account details
type Account struct {
	ID      int           `json:"id"`
	Name    string        `json:"name"`
	Balance float64       `json:"balance"`
	History []Transaction `json:"history"`
}

// Transaction details
type Transaction struct {
	Type      string    `json:"type"` // deposit / withdraw
	Amount    float64   `json:"amount"`
	Timestamp time.Time `json:"timestamp"`
}

// global variables
var accounts = make(map[int]*Account)
var nextID = 1
var mu sync.Mutex

// Create an account
func createAccount(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	var acc Account
	json.NewDecoder(r.Body).Decode(&acc)

	acc.ID = nextID
	acc.Balance = 0
	nextID++

	accounts[acc.ID] = &acc

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(acc)
}

// View Account
func getAccount(w http.ResponseWriter, r *http.Request) {
	// id := extractID(r.URL.Path)
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	acc, exists := accounts[id]
	if !exists {
		http.Error(w, "Account not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(acc)
}

// Deposit Money
func deposit(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	acc, exists := accounts[id]
	if !exists {
		http.Error(w, "Account not found", http.StatusNotFound)
		return
	}

	var req struct {
		Amount float64 `json:"amount"`
	}

	json.NewDecoder(r.Body).Decode(&req)

	if req.Amount <= 0 {
		http.Error(w, "Invalid amount", http.StatusBadRequest)
		return
	}

	acc.Balance += req.Amount

	acc.History = append(acc.History, Transaction{
		Type:      "deposit",
		Amount:    req.Amount,
		Timestamp: time.Now(),
	})

	json.NewEncoder(w).Encode(acc)
}

// Withdraw Money
func withdraw(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	acc, exists := accounts[id]
	if !exists {
		http.Error(w, "Account not found", http.StatusNotFound)
		return
	}

	var req struct {
		Amount float64 `json:"amount"`
	}

	json.NewDecoder(r.Body).Decode(&req)

	if req.Amount <= 0 {
		http.Error(w, "Invalid amount", http.StatusBadRequest)
		return
	}

	if acc.Balance < req.Amount {
		http.Error(w, "Insufficient funds", http.StatusBadRequest)
		return
	}

	acc.Balance -= req.Amount

	acc.History = append(acc.History, Transaction{
		Type:      "withdraw",
		Amount:    req.Amount,
		Timestamp: time.Now(),
	})

	json.NewEncoder(w).Encode(acc)
}

// View transactions
func getTransactions(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	acc, exists := accounts[id]
	if !exists {
		http.Error(w, "Account not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(acc.History)
}

// Extract id from url
func extractID(path string) int {
	parts := strings.Split(path, "/")
	id, _ := strconv.Atoi(parts[2])
	return id
}
