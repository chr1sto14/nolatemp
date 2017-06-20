package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var Db *sql.DB = openDb()

func openDb() *sql.DB {
	dbhostname := mustGetenv("RDS_HOSTNAME")
	dbport := mustGetenv("RDS_PORT")
	dbname := mustGetenv("RDS_DB_NAME")
	user := mustGetenv("RDS_USERNAME")
	password := mustGetenv("RDS_PASSWORD")

	db, err := sql.Open("postgres",
		fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s",
			dbhostname, dbport, dbname, user, password))
	if err != nil {
		log.Printf("error connecting to the database: %v", err)
	}
	return db
}

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Printf("%s environment variable not set.", k)
	}
	return v
}
