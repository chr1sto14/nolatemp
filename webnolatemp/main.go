package main

import (
	"fmt"
	"github.com/chr1sto14/nolatemp/temp"
	"net/http"
)

func tempHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Temperature coming right up!")
}

func nolaHandler(w http.ResponseWriter, r *http.Request) {
	outsideTemp := temp.GetNolaTemp()
	fmt.Printf("The temperature in NOLA is %g F.", outsideTemp)
}

func main() {
	http.HandleFunc("/temp", tempHandler)
	http.HandleFunc("/nola", nolaHandler)
	http.ListenAndServe(":8080", nil)

	// TODO
	// 1. receieve data from rpi
	// 3. store data (ts, inside, outside) to cockroachdb
	// 4. recieve commands from hipchat
	// 5. gather data from db based upon timeline
	// 6. format a plot ( inside, outside vs. time )
	// 7. format response to nola hipchat
}
