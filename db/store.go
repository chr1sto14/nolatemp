package db

import (
	"time"

	_ "github.com/lib/pq"
)

func InsertTsInOut(ts time.Time, inTemp float64, outTemp float64) error {
	_, err := Db.Exec(
		"INSERT INTO nolatemp.temp (ts, intemp, outtemp) VALUES ($1, $2, $3)",
		ts,
		inTemp,
		outTemp,
	)
	return err
}
