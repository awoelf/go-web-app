package services

import (
	"database/sql"
	"time"
)

var db *sql.DB
const dbTimeout = time.Second * 3

type Models struct {
	Comment Comment
}

func Register(dbPool *sql.DB) Models {
	db = dbPool
	return Models{}
}