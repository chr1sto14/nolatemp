package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DbUrl string
var Db *sql.DB

func OpenDb() {
	db, err := sql.Open("postgres", DbUrl)
	if err != nil {
		log.Printf("error connecting to the database: %v", err)
	}
	Db = db
	return
}
