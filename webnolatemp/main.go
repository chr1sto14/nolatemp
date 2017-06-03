package main

import (
	"fmt"
	"github.com/chr1sto14/nolatemp/db"
	"github.com/chr1sto14/nolatemp/net"
	"github.com/chr1sto14/nolatemp/temp"
	"log"
	"net/http"
)

func tempHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Temperature coming right up!")
}

func nolaHandler(w http.ResponseWriter, r *http.Request) {
	// parse json
	// should just return ts and inTemp
	ts, inTemp, err := temp.ParseTemp(r.Body)
	if err != nil {
		net.Bad(w)
		defer r.Body.Close()
		log.Printf("Error: %v", err)
		return
	}
	defer r.Body.Close()

	// prepare ts, inTemp, outTemp
	// TODO err
	outTemp := temp.GetNolaTemp()
	log.Printf("The ts %v", ts)
	log.Printf("The inTemp %v", inTemp)
	log.Printf("The outTemp %v", outTemp)

	// cockroach db
	err = db.InsertTsInOut(ts, inTemp, outTemp)
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC) // only show UTC time
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
