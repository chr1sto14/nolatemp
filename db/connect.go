package db

import (
	"bytes"
	"database/sql"
	"io/ioutil"
	"log"
	"os/user"

	_ "github.com/lib/pq"
)

var Db *sql.DB = openDb()

func openDb() *sql.DB {
	usr, _ := user.Current()

	datab, _ := ioutil.ReadFile(usr.HomeDir + "/nolatemp.properties")
	dbUrl := string(bytes.TrimSpace(datab))

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}
	return db
}
