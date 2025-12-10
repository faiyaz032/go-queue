package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/faiyaz032/go-queue/internal/jobqueue"
	redisclient "github.com/faiyaz032/go-queue/internal/redis"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	rdb := redisclient.New("localhost:6379")
	processor := func(job jobqueue.Job) {
		fmt.Println("Processing job like sending email:", job)
		time.Sleep(1 * time.Second)
	}
	queue := jobqueue.NewQueue(rdb, 5, processor)
	queue.Run(ctx)

	for i := 1; i <= 5; i++ {
		job := jobqueue.Job{
			ID:   fmt.Sprintf("%d", i),
			Type: "email",
			Data: map[string]string{"to": "faiyaz@example.com"},
		}
		queue.Enqueue(job)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	fmt.Println("Shutting down...")
	cancel()
	time.Sleep(1 * time.Second)
}
