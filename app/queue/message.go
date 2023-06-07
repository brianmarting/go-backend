package queue

type Message interface {
	GetBytes() []byte
	Ack()
}
