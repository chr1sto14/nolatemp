package main

import (
	"github.com/chr1sto14/nolatemp/temp"
	"log"
	"os"
	"time"
)

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

	// check env variable for inittemp
	err = temp.InitTemp()
	if err != nil {
		log.Printf("Error: %v", err)
	}

	// get timestamp
	now, err := time.Now().MarshalJSON()
	if err != nil {
		log.Printf("Error: %v", err)
	}

	// read temperature from rpi
	// TODO temp.GetTemp()
	tempVal := float64(100)

	err = temp.SendTemp(now, tempVal)
	if err != nil {
		log.Printf("Error: %v", err)
	}
}
