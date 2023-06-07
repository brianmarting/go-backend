package socket

type Listener interface {
	Start() (<-chan Message, chan<- string, error)
}
