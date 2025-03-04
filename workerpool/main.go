package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	ID       int
	RandomNo int
}

type Result struct {
	JobID  int
	Output int
}

func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job.ID)
		time.Sleep(time.Second)
		results <- Result{JobID: job.ID, Output: job.RandomNo * 2}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	numWorkers := 3
	numJobs := 10

	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)

	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	for i := 1; i <= numJobs; i++ {
		jobs <- Job{ID: i, RandomNo: rand.Intn(100)}
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Printf("Job ID %d processed with output %d\n", result.JobID, result.Output)
	}
}
