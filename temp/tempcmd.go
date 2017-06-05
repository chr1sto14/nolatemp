package temp

import (
	"errors"
	"github.com/chr1sto14/nolatemp/db"
	"github.com/chr1sto14/nolatemp/hipchat"
	"github.com/chr1sto14/nolatemp/tempplot"
	"log"
	"strconv"
	"strings"
	"time"
)

func ResolveCmd(cmd string) (timeType string, help bool, err error) {
	// cmd should be of the form "/temp _____"
	z := strings.SplitN(cmd, " ", 2)
	if len(z) == 0 {
		err = errors.New("No command received.")
		return
	} else if len(z) == 1 {
		if z[0] == "/temp" {
			help = true
			return
		} else {
			err = errors.New("Improper command.")
			help = true
			return
		}
	}
	timeType = z[1]
	return
}

func tempNow() (rv interface{}, err error) {
	// 1. get latest reading from db
	// 2. format into hipchat message
	return
}

// 2. get latest readings from db based upon time
// 3. form plot
// 4. save plot in db
// 5. return formatted hipchat with link to img in db
func tempTime(t string) (rv interface{}, err error) {
	loc, err := time.LoadLocation("America/Chicago")
	now := time.Now().In(loc)
	hour := 1
	switch t {
	case "year":
		hour *= 12
		fallthrough
	case "month":
		hour *= 4
		fallthrough
	case "week":
		hour *= 7
		fallthrough
	case "day":
		hour *= 24
		fallthrough
	default:
		hour *= 1
	}

	// TODO get data from db
	cutoff := now.Add(time.Duration(-hour) * time.Hour)
	log.Printf("time: %v", cutoff)
	tss, intemps, outtemps, err := db.QueryTemp(cutoff)
	if err != nil {
		return
	}

	// TODO hand it off to a plotting package
	img, err := tempplot.MakePlot(tss, intemps, outtemps)
	if err != nil {
		return
	}

	// TODO store plot in db
	id, err := db.InsertImg(img.Bytes())
	if err != nil {
		return
	}

	// TODO format hipchat message
	rv = hipchat.MsgImgUrl(strconv.FormatInt(id, 10))
	return
}

func DoCmd(timeType string) (interface{}, error) {
	if timeType == "now" {
		return tempNow()
	}
	return tempTime(timeType)
}
