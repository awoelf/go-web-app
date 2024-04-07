package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	Client *sql.DB
}

func Connect() (*Database, error) {
	// Open database
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// Check connection to db
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return &Database{
		Client: db,
	}, nil
}