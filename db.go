package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

// Go does NOT directly talk to MySQL/PostgreSQL.
// database/sql = abstraction layer (connection pool, query handling)
// Driver = actual implementation (MySQL, PostgreSQL, etc.)

type dbUser struct {
	ID   int
	Name string
}

type Result struct {
	UserName string `db:"user_name"`
	OrderID  int    `db:"order_id"`
}

func everythingAboutDb() {

	// Connecting to Database
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/dbname")

	// To verify connection:
	err = db.Ping()
	if err != nil {
		fmt.Println("Connection not established")
	}

	// manages connections
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Hour)

	// Executing Queries: INSERT / UPDATE / DELETE
	result, err := db.Exec("INSERT INTO users(name) VALUES(?)", "Kartik")

	// Get affected rows:
	rows1, _ := result.RowsAffected()
	fmt.Println("Number of rows:", rows1)

	// SELECT (Single Row)
	row := db.QueryRow("SELECT name FROM users WHERE id=?", 1)
	var name string
	err = row.Scan(&name)

	// SELECT (Multiple Rows)
	rows, err := db.Query("SELECT id, name FROM users")
	//defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)
	}

	// Struct Mapping (Manual ORM style)
	rows, _ = db.Query("SELECT id, name FROM users")

	var users []dbUser

	for rows.Next() {
		var u dbUser
		rows.Scan(&u.ID, &u.Name)
		users = append(users, u)
	}

	// Prepared Statements (Performance + Security)
	stmt, err := db.Prepare("INSERT INTO users(name) VALUES(?)")
	if err != nil {
		return
	}
	defer stmt.Close()
	stmt.Exec("Kartik")

	// Transactions
	tx, err := db.Begin()

	_, err = tx.Exec("UPDATE accounts SET balance = balance - 100 WHERE id=1")
	_, err = tx.Exec("UPDATE accounts SET balance = balance + 100 WHERE id=2")

	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
}

func sqlxDb() {

	// connectiong to database: this does both open and ping
	db, err := sqlx.Connect("mysql", "user:pass@tcp(localhost:3306)/db")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Fetch multiple rows
	var users []dbUser
	err = db.Select(&users, "SELECT id, name FROM users")

	// Fetch single row:
	var user dbUser
	err = db.Get(&user, "SELECT * FROM users WHERE id=?", 1)

	// Named quesries
	query := `
	INSERT INTO users (name, age)
	VALUES (:name, :age)
	`
	db.NamedExec(query, user)

	// The IN (?)
	var ids []int
	db.Select(&users, "SELECT * FROM users WHERE id IN (?)", ids)

	query, args, err := sqlx.In(
		"SELECT * FROM users WHERE id IN (?)",
		ids,
	)
	query = db.Rebind(query)
	err = db.Select(&users, query, args...)

	// Transactions
	tx, err := db.Beginx()

	_, err = tx.Exec("INSERT INTO users(name) VALUES(?)", "Kartik")

	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}

	// Struct + Joins Mapping
	query = `
	SELECT u.name AS user_name, o.id AS order_id
	FROM users u
	JOIN orders o ON u.id = o.user_id
	`
	var results []Result
	db.Select(&results, query)
}
