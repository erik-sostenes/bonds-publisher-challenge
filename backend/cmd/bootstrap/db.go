package bootstrap

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

// Type represents an uint for the type of DataBase
type Type uint

const (
	// SQL represents MySQL database
	SQL Type = iota
	// NoSQL represents (MongoDB, Redis) database
	NoSQL
)

// Configuration represents the settings of the type of storage
type Configuration struct {
	// Type defines the type of storage to be used.
	Type
	DSN string
}

// NewPostgresClient method that will connect a postgres client and returns an instance of postgres.Client
func NewPostgresClient(config Configuration) (*sql.DB, error) {
	if config.DSN == "" {
		panic("missing environment variable 'DSN_POSTGRESQL'")
	}

	switch config.Type {
	case SQL:
		client, err := sql.Open("postgres", config.DSN)
		if err != nil {
			return client, err
		}

		return client, client.Ping()
	default:
		panic(fmt.Sprintf("%T type is not supported", config.Type))
	}
}

func PostgreSQLInjector() (db *sql.DB) {
	config := Configuration{
		Type: SQL,
		DSN:  strings.TrimSpace(os.Getenv("DSN_POSTGRESQL")),
	}

	db, err := NewPostgresClient(config)
	if err != nil {
		panic(err)
	}
	return
}
