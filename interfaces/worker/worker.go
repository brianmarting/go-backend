package worker

type WorkerJob interface {
	Work() error
}
