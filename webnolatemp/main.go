package main

import (
	"github.com/chr1sto14/nolatemp/db"
	"github.com/chr1sto14/nolatemp/hipchat"
	"github.com/chr1sto14/nolatemp/net"
	"github.com/chr1sto14/nolatemp/temp"
	"log"
	"net/http"
	"path"
)

func cmdHandler(w http.ResponseWriter, r *http.Request) {
	cmdJson, err := hipchat.ParseCmd(r.Body)
	if err != nil {
		net.Bad(w) // TODO following example http wiki example
		log.Printf("Error: %v", err)
		return
	}

	timeType, help, err := temp.ResolveCmd(cmdJson.Item.Message)
	if err != nil {
		net.Bad(w)
		log.Printf("Error: %v", err)
		return
	} else if help {
		net.Json(w, hipchat.Help())
		return
	}

	val, err := temp.DoCmd(timeType)
	if err != nil {
		net.Bad(w)
		log.Printf("Error: %v", err)
		return
	}
	net.Json(w, val)
}

func imgHandler(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)

	log.Printf("id: %v", id)
	buf, err := db.QueryImg(id)
	if err != nil {
		net.Bad(w)
		log.Printf("Error: %v", err)
		return
	}
	net.Img(w, buf)
}

func nolaHandler(w http.ResponseWriter, r *http.Request) {
	// parse and return ts and inTemp
	ts, inTemp, err := temp.ParseTemp(r.Body)
	if err != nil {
		net.Bad(w)
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
	if err != nil {
		net.Bad(w)
		log.Printf("Error: %v", err)
		return
	}
	net.Good(w)
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC) // only show UTC time
	http.HandleFunc("/temp", cmdHandler)
	http.HandleFunc("/img/", imgHandler)
	http.HandleFunc("/nola", nolaHandler)
	http.ListenAndServe(":8888", nil)

	// TODO
	// 1. receive cmds from hipchat
	// 2. get relevant data from db
	// 3. store image to db
	// 4. recieve commands from hipchat
	// 5. gather data from db based upon timeline
	// 6. format a plot ( inside, outside vs. time )
	// 7. format response to nola hipchat
}
