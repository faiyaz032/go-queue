package jobqueue

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Queue struct {
	RDB       *redis.Client
	JobChan   chan Job
	WorkerNum int
	Processor func(Job)
}

func NewQueue(rdb *redis.Client, workerNum int, processor func(Job)) *Queue {
	return &Queue{
		RDB:       rdb,
		JobChan:   make(chan Job),
		WorkerNum: workerNum,
		Processor: processor,
	}
}

func (q *Queue) Run(ctx context.Context) {
	go q.dispatcher(ctx)

	for i := 1; i < q.WorkerNum; i++ {
		go q.worker(i, ctx)
	}
}

func (q *Queue) Enqueue(job Job) error {
	return q.RDB.RPush(context.Background(), "job_queue", job.ToJSON()).Err()
}
