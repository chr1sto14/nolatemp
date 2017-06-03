package temp

import (
	"errors"
	"github.com/chr1sto14/nolatemp/net"
	"io"
	"time"
)

// TODO make this a cmd line arg?
var webnolatempurl string = "http://localhost:8888/nola"

type TempJson struct {
	Ts   []byte  `json:"ts"`
	Temp float64 `json:"temp"`
}

func SendTemp(ts []byte, temp float64) error {
	msg := TempJson{ts, temp}
	return net.SendJson(webnolatempurl, msg)
}

func ParseTemp(body io.Reader) (ts time.Time, inTemp float64, err error) {
	var tempJson TempJson
	err = net.ParseJson(body, &tempJson)
	if err != nil {
		return
	}

	if tempJson.Ts == nil || tempJson.Temp == 0 {
		err = errors.New("Ts or Temp not present in json.")
		return
	}

	// TODO err
	err = ts.UnmarshalJSON(tempJson.Ts)
	inTemp = tempJson.Temp
	return
}
