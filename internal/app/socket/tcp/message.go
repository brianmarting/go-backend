package tcp

type Message struct {
	body []byte
}

func (m Message) GetBody() []byte {
	return m.body
}
