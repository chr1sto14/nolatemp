package db

import "time"

func InsertTsInOut(ts time.Time, inTemp float64, outTemp float64) (err error) {
	_, err = Db.Exec(
		"INSERT INTO nolatemp.temp (ts, intemp, outtemp) VALUES ($1, $2, $3)",
		ts,
		inTemp,
		outTemp,
	)
	return
}

func InsertImg(img []byte) (id int64, err error) {
	err = Db.QueryRow(
		"INSERT INTO nolatemp.images (img) VALUES ($1) "+
			"RETURNING id",
		img).Scan(&id)
	return
}
