package main

import (
	"flag"
	"fmt"
	"github.com/chr1sto14/nolatemp/db"
	"github.com/chr1sto14/nolatemp/hipchat"
	"github.com/chr1sto14/nolatemp/net"
	"github.com/chr1sto14/nolatemp/temp"
	"log"
	"net/http"
	"os"
	"path"
)

func cmdHandler(w http.ResponseWriter, r *http.Request) {
	cmdJson, err := hipchat.ParseCmd(r.Body)
	if err != nil {
		net.Bad(w, err)
		log.Printf("Error: %v", err)
		return
	}

	timeType, help, err := temp.ResolveCmd(cmdJson.Item.Message)
	if err != nil {
		net.Bad(w, err)
		log.Printf("Error: %v", err)
		return
	} else if help {
		net.Json(w, hipchat.Help())
		return
	}

	val, err := temp.DoCmd(timeType)
	if err != nil {
		net.Bad(w, err)
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
		net.Bad(w, err)
		log.Printf("Error: %v", err)
		return
	}
	net.Img(w, buf)
}

func nolaHandler(w http.ResponseWriter, r *http.Request) {
	// parse and return ts and inTemp
	ts, inTemp, err := temp.ParseTemp(r.Body)
	if err != nil {
		net.Bad(w, err)
		log.Printf("Error: %v", err)
		return
	}
	defer r.Body.Close()

	// prepare ts, inTemp, outTemp
	outTemp, err := temp.GetNolaTemp()
	if err != nil {
		net.Bad(w, err)
		log.Printf("Error: %v", err)
		return
	}

	// cockroach db
	err = db.InsertTsInOut(ts, inTemp, outTemp)
	if err != nil {
		net.Bad(w, err)
		log.Printf("Error: %v", err)
		return
	}
	net.Good(w)
}

func logSetup() (err error) {
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC) // only show UTC time
	f, err := os.OpenFile("nolatemp.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	log.SetOutput(f)
	return
}

func main() {
	err := logSetup()
	if err != nil {
		log.Printf("Error: %v", err)
	}

	// get command line args
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] [URL]\n\n", os.Args[0])
		fmt.Fprint(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}

	url := flag.String("url", "", "url for own server")

	if len(os.Args[1:]) == 0 {
		flag.Usage()
		os.Exit(0)
	}
	flag.Parse()
	if *url == "" {
		log.Printf("Error: url is required")
	}
	hipchat.MyUrl = *url

	http.HandleFunc("/temp", cmdHandler)
	http.HandleFunc("/img/", imgHandler)
	http.HandleFunc("/nola", nolaHandler)
	http.ListenAndServe(":8888", nil)
}
