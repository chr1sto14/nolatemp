package hipchat

import (
	"github.com/chr1sto14/nolatemp/net"
	"io"
)

type Msg2 struct {
	Message string `json:"message"`
}

type Msg1 struct {
	Message Msg2 `json:"message"`
}

type CmdJson struct {
	Item Msg1 `json:"item"`
}

func ParseCmd(body io.Reader) (cmdJson CmdJson, err error) {
	err = net.ParseJson(body, &cmdJson)
	return
}
