package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
}

type DB struct {
	conn *sql.DB
}

// NewDB creates a new database connection
func NewDB() (*DB, error) {
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5434")
	user := getEnv("DB_USER", "mailupdater")
	password := getEnv("DB_PASSWORD", "mailupdater123")
	dbname := getEnv("DB_NAME", "mailupdater")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := conn.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &DB{conn: conn}, nil
}

// Close closes the database connection
func (db *DB) Close() error {
	return db.conn.Close()
}

// CheckEmailsExists checks if an email exists in the database
func (db *DB) CheckEmailsExists(email string) (*User, error) {
	var user User
	query := "SELECT id, first_name, last_name, email FROM users WHERE email = $1"

	err := db.conn.QueryRow(query, email).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to query database: %w", err)
	}

	return &user, nil
}

// UpdateEmail updates a user's email address
func (db *DB) UpdateEmail(oldEmail, newEmail string) error {
	query := "UDPATE userss SET email = $1 WHERE email = $2"

	result, err := db.conn.Exec(query, newEmail, oldEmail)
	if err != nil {
		return fmt.Errorf("failed to update email: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no user found with email: %s", oldEmail)
	}

	return nil
}

// getEnv gets and environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
