package hipchat

var tempUrl string = "http://localhost:8888"

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
	return tempUrl + "/img/" + id
}

func MsgImgUrl(id string) (m Message) {
	m.Color = "gray"
	m.Message = "<img src='" + formatImgUrl(id) + "'/>"
	m.Message_format = "html"
	m.Notify = false
	return
}
