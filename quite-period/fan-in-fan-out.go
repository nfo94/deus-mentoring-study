package main

import (
	"fmt"
	"sync"
	"time"
)

// Fan-Out/Fan-In pattern

// The worker receives the id and the job
func workerFofi(id int, job int) int {
	time.Sleep(time.Millisecond * 100) // Simulate work
	fmt.Printf("Processing id %d\n", id)
	return job * job // Example operation
}

func runWorkerFofi() {
	// Slice of jobs that we'll work on
	jobs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Channel of results of the jobs
	results := make(chan int, len(jobs))
	// We're using wait groups to synchronize the work
	var wg sync.WaitGroup

	// Lauching go routines for each job
	for i, job := range jobs {
		// One go routine to keep track
		wg.Add(1)
		go func(id, job int) {
			// Mark as done, minus one go routine to keep track
			defer wg.Done()
			// Do the job
			result := workerFofi(id, job)
			// Throw the result int channel
			results <- result
		}(i, job)
	}

	// Wait for go routines to finish
	go func() {
		wg.Wait()
		// Close results channel when done
		close(results)
	}()

	sum := 0
	for result := range results {
		sum += result // Aggregate result
	}

	fmt.Printf("Sum of squares: %d\n", sum)
}
