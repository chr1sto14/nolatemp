package hipchat

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
