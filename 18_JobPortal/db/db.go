package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

var DB *sql.DB // database

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "./data/api.db")
	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTable()
}

func createTable() {
	createJobTable := `CREATE TABLE IF NOT EXISTS jobs (
	 id INTEGER PRIMARY KEY AUTOINCREMENT,
     title TEXT NOT NULL,
     description TEXT NOT NULL,
     skills TEXT NOT NULL,
	 experience INTEGER NOT NULL, 
	 salary_range TEXT NOT NULL,
	 location TEXT NOT NULL,
	 dateTime DATETIME NOT NULL,
	 user_id INTEGER NOT NULL,
     created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)`

	_, err := DB.Exec(createJobTable)

	if err != nil {
		panic("Error creating jobs table.")
	}

	createUserTable := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL,
		password TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	   )`

	_, err = DB.Exec(createUserTable)

	if err != nil {
		panic("Error creating user table.")
	}

}
