package main

import (
	"fmt"
	"time"
)

// Worker pool pattern

// Function `worker` that receives
// - An id of type int
// - A receiver-only `tasks` int channel
// - A sender-only `results` int channel
func workerWp(id int, tasks <-chan int, results chan<- int) {
	// In the range of tasks received in the channel
	for task := range tasks {
		fmt.Printf("Worker %d processing task %d\n", id, task)
		// Simulate some work time
		time.Sleep(time.Second)
		// Send the example operation result in the sender channel
		results <- task * task
	}
}

func runWorkerWp() {
	// These are our workers (go routines)
	numWorkers := 3
	// These are our tasks. They'll be sent to a channel
	numTasks := 5

	// The channel for the results
	results := make(chan int, numTasks)
	// The channel to send the tasks
	tasks := make(chan int, numTasks)

	// Starting our workers
	for i := 1; i <= numWorkers; i++ {
		go workerWp(i, tasks, results)
	}

	// Sending tasks for workers
	for j := 1; j <= numTasks; j++ {
		tasks <- j
	}
	close(tasks) // Close the task channel when it's done

	// Looping through results
	for k := 1; k <= numTasks; k++ {
		result := <-results
		fmt.Printf("Result: %d\n", result)
	}
}
