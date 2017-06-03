package main

import (
	"fmt"
	"github.com/chr1sto14/nolatemp/db"
	"github.com/chr1sto14/nolatemp/net"
	"github.com/chr1sto14/nolatemp/temp"
	"log"
	"net/http"
	"time"
)

func tempHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Temperature coming right up!")
}

func nolaHandler(w http.ResponseWriter, r *http.Request) {
	// parse json
	var tempJson temp.TempJson
	err := net.ParseJson(r.Body, &tempJson)
	if err != nil {
		// TODO  just print error
		panic(err)
	}
	defer r.Body.Close()
	log.Printf("The response %+v", tempJson)

	// prepare ts, inTemp, outTemp
	var ts time.Time
	// TODO err
	err = ts.UnmarshalJSON(tempJson.Ts)
	inTemp := tempJson.Temp
	outTemp := temp.GetNolaTemp()
	log.Printf("The ts %v", ts)
	log.Printf("The inTemp %v", inTemp)
	log.Printf("The outTemp %v", outTemp)

	// cockroach db
	err = db.InsertTsInOut(ts, inTemp, outTemp)
}

func main() {
	http.HandleFunc("/temp", tempHandler)
	http.HandleFunc("/nola", nolaHandler)
	http.ListenAndServe(":8888", nil)

	// TODO
	// 3. send success back from nolaHandler
	// 3. b) send failed back if not proper json complete
	// 4. recieve commands from hipchat
	// 5. gather data from db based upon timeline
	// 6. format a plot ( inside, outside vs. time )
	// 7. format response to nola hipchat
}
