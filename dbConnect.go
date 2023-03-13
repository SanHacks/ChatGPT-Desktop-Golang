package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

// Message represents a message in the chat
type Message struct {
	ID        int
	Sender    string
	Content   string
	CreatedAt time.Time
}

// addMessage adds a message to the database
func addMessage(sender string, content string) error {
	// Open a connection to the database
	db, err := sql.Open("sqlite3", "./messages.db")
	if err != nil {
		return err
	}
	defer db.Close()

	// Prepare a SQL statement to insert the message into the database
	stmt, err := db.Prepare("INSERT INTO messages (sender, content) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the prepared statement with the message as parameters
	_, err = stmt.Exec(sender, content)
	if err != nil {
		return err
	}

	return nil
}

// getMessages retrieves all messages from the database
func getMessages() ([]Message, error) {
	// Open a connection to the database
	db, err := sql.Open("sqlite3", "./messages.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Execute a SQL query to retrieve all messages
	rows, err := db.Query("SELECT id, sender, content, created_at FROM messages ORDER BY created_at ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over the result set and create a slice of Message structs

	var messages []Message
	for rows.Next() {
		var m Message
		if err := rows.Scan(&m.ID, &m.Sender, &m.Content, &m.CreatedAt); err != nil {
			return nil, err
		}
		messages = append(messages, m)
		log.Printf("Message: %v", m)

	}
	return messages, nil
}
