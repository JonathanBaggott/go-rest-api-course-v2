package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Database is a struct that represents a database connection.
type Database struct {
	Client *sqlx.DB
}

// NewDatabase creates a new Database instance and establishes a connection to the database.
func NewDatabase() (*Database, error) {
	// Create a connection string using environment variables.
	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_TABLE"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("SSL_MODE"),
	)
	// Connect to the PostgreSQL database using the sqlx package.
	dbConn, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		return &Database{}, fmt.Errorf("could not connect to the database: %w", err)
	}

	return &Database{
		Client: dbConn,
	}, nil
}

// Ping sends a ping request to the database server to check if the connection is alive.
func (d *Database) Ping(ctx context.Context) error {
	return d.Client.DB.PingContext(ctx)
}
