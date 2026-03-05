package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Account struct {
	ID      int     `json:"id" db:"id"` // changed to int for AUTO_INCREMENT
	Name    string  `json:"name" db:"name"`
	Balance float64 `json:"balance" db:"balance"`
}

type Transaction struct {
	ID        int     `json:"id"`
	AccountID int     `json:"account_id"` // int for foreign key
	Type      string  `json:"type"`
	Amount    float64 `json:"amount"`
}

func main() {
	var err error
	// MySQL connection string: username:password@tcp(host:port)/dbname
	//db, err = sql.Open("mysql", "root:2506sql(127.0.0.1:3306)/bank")
	db, err = sql.Open("mysql", "root:2506sql@tcp(127.0.0.1:3306)/bank")
	// db is usually a variable of type *sql.DB
	if err != nil {
		log.Fatal(err)
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	createTables()

	http.HandleFunc("/create", createAccount)
	http.HandleFunc("/account", getAccount)
	http.HandleFunc("/deposit", deposit)
	http.HandleFunc("/withdraw", withdraw)
	http.HandleFunc("/transactions", getTransactions)

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}

func createTables() {
	// Create accounts table
	accountsTable := `
	CREATE TABLE IF NOT EXISTS accounts (
		id INT PRIMARY KEY AUTO_INCREMENT,
		name VARCHAR(100),
		balance DOUBLE
	);`
	_, err := db.Exec(accountsTable)
	if err != nil {
		log.Fatal("Error creating accounts table: ", err)
	}

	// Create transactions table
	transactionsTable := `
	CREATE TABLE IF NOT EXISTS transactions (
		id INT PRIMARY KEY AUTO_INCREMENT,
		account_id INT,
		type VARCHAR(20),
		amount DOUBLE,
		FOREIGN KEY (account_id) REFERENCES accounts(id)
	);`
	_, err = db.Exec(transactionsTable)
	if err != nil {
		log.Fatal("Error creating transactions table: ", err)
	}
}

// Create Account
func createAccount(w http.ResponseWriter, r *http.Request) {
	var acc Account
	json.NewDecoder(r.Body).Decode(&acc)

	res, err := db.Exec("INSERT INTO accounts(name,balance) VALUES(?,?)",
		acc.Name, acc.Balance)

	if err != nil {
		fmt.Println("Unable to insert into account")
		http.Error(w, err.Error(), 400)
		return
	}

	id, _ := res.LastInsertId()
	acc.ID = int(id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(acc)
}

// Deposit
func deposit(w http.ResponseWriter, r *http.Request) {
	var input struct {
		AccountID int     `json:"account_id"`
		Amount    float64 `json:"amount"`
	}
	json.NewDecoder(r.Body).Decode(&input)

	tx, _ := db.Begin()

	_, err := tx.Exec("UPDATE accounts SET balance = balance + ? WHERE id = ?",
		input.Amount, input.AccountID)

	if err != nil {
		tx.Rollback()
		http.Error(w, err.Error(), 400)
		return
	}

	_, err = tx.Exec("INSERT INTO transactions(account_id,type,amount) VALUES(?,?,?)",
		input.AccountID, "deposit", input.Amount)

	if err != nil {
		tx.Rollback()
		http.Error(w, err.Error(), 400)
		return
	}

	tx.Commit()
	w.Write([]byte("Deposit Successful"))
}

// Withdraw
func withdraw(w http.ResponseWriter, r *http.Request) {
	var input struct {
		AccountID int     `json:"account_id"`
		Amount    float64 `json:"amount"`
	}
	json.NewDecoder(r.Body).Decode(&input)

	var balance float64
	db.QueryRow("SELECT balance FROM accounts WHERE id=?", input.AccountID).Scan(&balance)

	if balance < input.Amount {
		http.Error(w, "Insufficient funds", 400)
		return
	}

	tx, _ := db.Begin()

	_, err := tx.Exec("UPDATE accounts SET balance = balance - ? WHERE id=?",
		input.Amount, input.AccountID)
	if err != nil {
		tx.Rollback()
		http.Error(w, err.Error(), 400)
		return
	}

	_, err = tx.Exec("INSERT INTO transactions(account_id,type,amount) VALUES(?,?,?)",
		input.AccountID, "withdraw", input.Amount)
	if err != nil {
		tx.Rollback()
		http.Error(w, err.Error(), 400)
		return
	}

	tx.Commit()
	w.Write([]byte("Withdraw Successful"))
}

// View Account
func getAccount(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var acc Account
	err := db.QueryRow("SELECT id,name,balance FROM accounts WHERE id=?", id).
		Scan(&acc.ID, &acc.Name, &acc.Balance)

	if err != nil {
		http.Error(w, "Not found", 404)
		return
	}

	json.NewEncoder(w).Encode(acc)
}

// View Transactions
func getTransactions(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	rows, _ := db.Query("SELECT id,account_id,type,amount FROM transactions WHERE account_id=?", id)
	defer rows.Close()

	var txs []Transaction

	for rows.Next() {
		var t Transaction
		rows.Scan(&t.ID, &t.AccountID, &t.Type, &t.Amount)
		txs = append(txs, t)
	}

	json.NewEncoder(w).Encode(txs)
}
