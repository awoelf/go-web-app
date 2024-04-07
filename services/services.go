package services

import (
	"database/sql"
	"log"
)

var db *sql.DB

type Models struct {

}

func Register(dbPool *sql.DB) Models {
	db = dbPool
	log.Print(db)
	return Models{}
}