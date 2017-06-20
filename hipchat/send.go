package hipchat

import (
	"fmt"
	"os"
)

var tempNowStr string = "<p><b>Time: </b>%s</p><p><b>Inside Temp: </b>%s&#8457</p><p><b>Outside Temp: </b>%s&#8457</p>"

var MyUrl string = os.Getenv("MYURL")

type Message struct {
	Color          string `json:"color"`
	Message        string `json:"message"`
	Message_format string `json:"message_format"`
	Notify         bool   `json:"notify"`
}

func Help() (m Message) {
	m.Color = "green"
	m.Message = "Usage: /temp [now|hour|day|week|month|year]"
	m.Message_format = "text"
	m.Notify = false
	return
}

func formatImgUrl(id string) string {
	return MyUrl + "/img/" + id
}

func MsgImgUrl(id string) (m Message) {
	m.Color = "gray"
	m.Message = "<img src='" + formatImgUrl(id) + "'/>"
	m.Message_format = "html"
	m.Notify = false
	return
}

func MsgTempNow(ts string, intemp string, outtemp string) (m Message) {
	m.Color = "gray"
	m.Message = fmt.Sprintf(tempNowStr, ts, intemp, outtemp)
	m.Message_format = "html"
	m.Notify = false
	return
}
