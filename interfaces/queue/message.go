package queue

type Message interface {
	GetBytes() []byte
}
