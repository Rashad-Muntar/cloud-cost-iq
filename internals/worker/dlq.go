package worker

type DLQ struct {
	failedJobs []Job
}

func NewDLQ() *DLQ {
	return &DLQ{
		failedJobs: make([]Job, 0),
	}
}

func (d *DLQ) Push(job Job) {
	d.failedJobs = append(d.failedJobs, job)
}