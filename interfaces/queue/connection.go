package queue

type Connection interface {
	GetConnection(url string) error
}
