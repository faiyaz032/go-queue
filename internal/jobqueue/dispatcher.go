package jobqueue

import (
	"context"
	"fmt"
)

func (q *Queue) dispatcher(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Dispatcher stopped")
			return
		default:
			result, err := q.RDB.BLPop(ctx, 0, "job_queue").Result()
			if err != nil {
				continue
			}
			job := FromJSON(result[1])
			q.JobChan <- job
		}
	}
}
