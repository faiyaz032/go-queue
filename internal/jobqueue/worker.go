package jobqueue

import (
	"context"
	"fmt"
)

func (q *Queue) worker(id int, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d stopped\n", id)
			return
		case job := <-q.JobChan:
			fmt.Printf("Worker %d processing job %s\n", id, job.ID)
			q.Processor(job)
			fmt.Printf("Worker %d finished job %s\n", id, job.ID)
		}
	}
}
