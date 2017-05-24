package main

import (
	"fmt"
	"net/http"
)

func tempHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Temperature coming right up!")
}

func main() {
	http.HandleFunc("/temp", handler)
	http.ListenAndServe(":8080", nil)

	// TODO
	// 1. receieve data from rpi
	// 2. store data (ts, inside, outside) to cockroachdb
	// 3. recieve commands from hipchat
	// 4. gather data from db based upon timeline
	// 5. format a plot ( inside, outside vs. time )
	// 6. format response to nola hipchat
}
