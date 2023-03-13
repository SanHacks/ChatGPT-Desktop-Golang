package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func createDatabase() error {
	// Open a connection to the database
	db, err := sql.Open("sqlite3", "./messages.db")
	if err != nil {
		return err
	}
	defer db.Close()

	//check if the database exists
	//if it does, return
	//if it doesn't, create it

	if db == nil {
		log.Println("Database does not exist")
	}
	// Execute the SQL command to create the table
	_, err = db.Exec(`
        CREATE TABLE messages (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            sender TEXT NOT NULL,
            content TEXT NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        )
    `)
	if err != nil {
		return err
	}
	log.Printf("Database created successfully")

	return nil
}
