package db

import (
	"time"

	_ "github.com/lib/pq"
)

type TempRow struct {
	Ts      time.Time `json:"ts"`
	InTemp  float64   `json:"intemp"`
	OutTemp float64   `json:"outtemp"`
}

func QueryTemp(minTs time.Time) (tss []time.Time, intemps []float64, outtemps []float64, err error) {
	rows, err := Db.Query(
		"SELECT ts, intemp, outtemp FROM nolatemp.temp WHERE ts > $1",
		minTs,
	)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var ts time.Time
		var intemp, outtemp float64
		if err = rows.Scan(&ts, &intemp, &outtemp); err != nil {
			return
		}
		tss = append(tss, ts)
		intemps = append(intemps, intemp)
		outtemps = append(outtemps, outtemp)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}

func QueryImg(id string) (buf []byte, err error) {
	err = Db.QueryRow("SELECT img FROM  nolatemp.images WHERE id = $1", id).Scan(&buf)
	return
}
