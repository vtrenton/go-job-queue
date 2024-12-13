package main

import (
	"fmt"
	"time"
)

type Job struct {
	ID string,
	Payload string,
	Status string
}

func main() {

}

// Create a queue for jobs to be stored
func AddJob(queue chan Job, job Job) {
	queue <- job
}

// Workers that spawn as go routines to manage and work on jobs.
func Worker(queue chan Job, workerID int) {
	for job := range queue {
		fmt.Printf("Worker %d processing job %s", workerID, job.ID)
		time.Sleep(1 * time.Second)
	}
}

// Spawn concurrent workers to execute based on the number of jobs in the queue
func StartDispatcher(queue chan Job, numWorkers int) {
	for i := 0; i < numWorkers; i++) {
		go Worker(queue, i)
	}
}
