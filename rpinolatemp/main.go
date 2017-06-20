package main

import (
	"flag"
	"fmt"
	"github.com/chr1sto14/nolatemp/temp"
	"github.com/chr1sto14/nolatemp/temprpi"
	"log"
	"os"
	"time"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC) // only show UTC time
	f, err := os.OpenFile("nolatemp.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("Error: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	// get command line args
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] [URL]\n\n", os.Args[0])
		fmt.Fprint(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}
	url := flag.String("url", "", "url for storing temperature")

	if len(os.Args[1:]) == 0 {
		flag.Usage()
		os.Exit(0)
	}
	flag.Parse()

	if *url == "" {
		log.Printf("Error: url is required")
	}

	// check env variable for inittemp
	err = temprpi.InitTemp()
	if err != nil {
		log.Printf("Error: %v", err)
	}

	// get timestamp
	now, err := time.Now().MarshalJSON()
	if err != nil {
		log.Printf("Error: %v", err)
	}

	// read temperature from rpi
	tempVal, err := temprpi.GetTemp()
	if err != nil {
		log.Printf("Error: %v", err)
	}

	err = temp.SendTemp(*url, now, tempVal)
	if err != nil {
		log.Printf("Error: %v", err)
	}
}
