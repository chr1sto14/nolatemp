package main

import (
	//"github.com/chr1sto14/nolatemp/temp"
	"log"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC) // only show UTC time

	// temp.GetTemp()
	temp := float64(100)

	SendTemp(temp)

	// TODO
	// 1. send data (ts, inside) to webnolatemp
}
