package main

import (
	"github.com/chr1sto14/nolatemp/temp"
	"log"
	"time"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC) // only show UTC time

	// get timestamp
	now, _ := time.Now().MarshalJSON()

	// read temperature from rpi
	// temp.GetTemp()
	tempVal := float64(100)

	err := temp.SendTemp(now, tempVal)
	if err != nil {
		panic(err)
	}

	// TODO
	// 1. send data (ts, inside) to webnolatemp
}
