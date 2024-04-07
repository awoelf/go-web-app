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

	createTable := `create table if not exists comments (
		id text not null primary key,
		name text not null,
		subject text not null,
		commentText text not null,
		createdAt datetime default current_timestamp,
		updatedAt datetime default current_timestamp
	)`

	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}
	
	// Check connection to db
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return &Database{
		Client: db,
	}, nil
}
