package main

import (
	"fmt"
	"time"
)

type Job struct {
	ID      string
	Payload string
	Status  string
}

func main() {
	jobQueue := make(chan Job, 10)
	StartDispatcher(jobQueue, 3)

	for i := 1; i <= 5; i++ {
		job := Job{
			ID:      fmt.Sprintf("job-%d", i),
			Payload: fmt.Sprintf("Payload for job-%d", i),
		}
		AddJob(jobQueue, job)
	}
	time.Sleep(5 * time.Second)
}

// Create a queue for jobs to be stored
func AddJob(queue chan Job, job Job) {
	queue <- job
}

// Workers that spawn as go routines to manage and work on jobs.
func Worker(queue chan Job, workerID int) {
	for job := range queue {
		fmt.Printf("Worker %d processing job %s\n", workerID, job.ID)
		time.Sleep(1 * time.Second)
	}
}

// Spawn concurrent workers to execute based on the number of jobs in the queue
func StartDispatcher(queue chan Job, numWorkers int) {
	for i := 0; i < numWorkers; i++ {
		go Worker(queue, i)
	}
}
