package main

import (
	"github.com/chr1sto14/nolatemp/db"
	"github.com/chr1sto14/nolatemp/hipchat"
	"github.com/chr1sto14/nolatemp/net"
	"github.com/chr1sto14/nolatemp/temp"
	"log"
	"net/http"
	"os"
	"path"
)

const indexPage = "public/index.html"
const favIcon = "public/favicon.ico"

func cmdHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		log.Printf("Body: %v", r.Body)
		cmdJson, err := hipchat.ParseCmd(r.Body)
		if err != nil {
			net.Bad(w, err)
			log.Printf("Error: %v", err)
			return
		}

		timeType, help, err := temp.ResolveCmd(cmdJson.Item.Message.Message)
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
}

func imgHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
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
}

func nolaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
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

func root(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.ServeFile(w, r, indexPage)
	}
}

func favicon(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.ServeFile(w, r, favIcon)
	}
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC) // only show UTC time
	f, err := os.OpenFile("/var/app/current/nolatemp.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("Error: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	http.HandleFunc("/", root)
	http.HandleFunc("/temp", cmdHandler)
	http.HandleFunc("/img/", imgHandler)
	http.HandleFunc("/nola", nolaHandler)
	http.HandleFunc("/favicon.ico", favicon)
	log.Printf("Listening on port %s\n\n", port)
	http.ListenAndServe(":"+port, nil)
}
