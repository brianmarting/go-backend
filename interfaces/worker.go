package interfaces

type WorkerJob interface {
	Work() error
}
