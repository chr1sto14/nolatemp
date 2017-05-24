package main

import (
	"github.com/chr1sto14/nolatemp/temp"
	"log"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC) // only show UTC time

	temp.GetTemp()

	// TODO
	// 1. get outside weather
	// 2. send data (ts, inside, outside) to webnolatemp
}
