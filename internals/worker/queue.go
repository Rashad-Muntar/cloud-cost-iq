package worker

type Queue struct {
	jobs chan Job
}

func NewQueue(buffer int) *Queue {
	return &Queue{
		jobs: make(chan Job, buffer),
	}
}

func (q *Queue) Push(job Job) {
	q.jobs <- job
}

func (q *Queue) Pop() <-chan Job {
	return q.jobs
}