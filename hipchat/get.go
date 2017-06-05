package hipchat

import (
	"github.com/chr1sto14/nolatemp/net"
	"io"
)

type Msg struct {
	Message string `json:"message"`
}

type CmdJson struct {
	Item Msg `json:"item"`
}

func ParseCmd(body io.Reader) (cmdJson CmdJson, err error) {
	err = net.ParseJson(body, &cmdJson)
	return
}
