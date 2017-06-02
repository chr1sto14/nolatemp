package main

import (
	"github.com/chr1sto14/nolatemp/net"
)

var webnolatempurl string = "http://localhost:8080/nola"

type TempJson struct {
	Temp float64 `json:"temp"`
}

func SendTemp(temp float64) error {
	msg := TempJson{temp}
	return net.SendJson(webnolatempurl, msg)
}
