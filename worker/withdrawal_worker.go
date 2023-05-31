package worker

type WithdrawalWorkerJob struct {
	WorkFn func() error
}

func (w *WithdrawalWorkerJob) Work() error {
	return w.WorkFn()
}
