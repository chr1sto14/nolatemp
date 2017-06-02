package temp

import (
	"github.com/chr1sto14/nolatemp/net"
)

// TODO make this a cmd line arg?
var webnolatempurl string = "http://localhost:8080/nola"

type TempJson struct {
	Ts   []byte  `json:"ts"`
	Temp float64 `json:"temp"`
}

func SendTemp(ts []byte, temp float64) error {
	msg := TempJson{ts, temp}
	return net.SendJson(webnolatempurl, msg)
}
