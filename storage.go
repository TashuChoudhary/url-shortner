package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("sqlite3", "./urls.db")
	if err != nil {
		log.Fatal("Failed to open database", err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS urls(
	code TEXT PRIMARY KEY,
	original_url TEXT NOT NULL
	);`
	if _, err := db.Exec(createTable); err != nil {
		log.Fatal("Failed to create table", err)

	}
}

func saveURL(code, originalURL string) error {
	_, err := db.Exec("INSERT INTO urls (code, original_url) VALUES (?, ?)", code, originalURL)
	return err
}

func getOriginalURL(code string) (string, bool) {
	row := db.QueryRow("SELECT original_url FROM urls WHERE code = ?", code)
	var originalURL string
	err := row.Scan(&originalURL)
	if err != nil {
		return "", false

	}
	return originalURL, true

}
