package worker

import "context"

type Pool struct {
	queue *Queue

	workers int

	processor *Processor
}

func NewPool(
	queue *Queue,
	workers int,
	processor *Processor,
) *Pool {

	return &Pool{
		queue: queue,
		workers: workers,
		processor: processor,
	}
}


func (p *Pool) Start(ctx context.Context) {

	for i := 0; i < p.workers; i++ {

		go func(workerID int) {

			for {
				select {

				case job := <-p.queue.Pop():

					_ = p.processor.Process(ctx, job)

				case <-ctx.Done():
					return
				}
			}

		}(i)
	}
}