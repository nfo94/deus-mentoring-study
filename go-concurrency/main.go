package main

// 1. Batch File Processor
// Goal: Practice using sync.WaitGroup to manage a batch of independent tasks.
// Problem: You have a list of 100 file paths (e.g., ["file_1.pdf", "file_2.pdf", ...]). You need to "process" all of them.
// Requirements:
// Write a main function that iterates through this list.
// For each file path, launch a new goroutine that calls a function processFile(filename string).
// Inside processFile, simulate work by sleeping for 1 second (time.Sleep(1 * time.Second)) and then printing "Processed [filename]".
// The main function must not exit until all 100 goroutines have finished.
// After all files are processed, main should print "All files processed. Shutting down."
// Primitive to use: sync.WaitGroup

// Solution
// func generateFilePath() [100]string {
// 	var arr [100]string
// 	for i := range arr {
// 		arr[i] = fmt.Sprintf("path%d", i+1)
// 	}
// 	return arr
// }

// func processFile(filepath string, wg *sync.WaitGroup) {
// 	defer wg.Done() // Mark as done, decreses the counter
// 	time.Sleep(1 * time.Second)
// 	fmt.Printf("Processed %s\n", filepath)
// }

// func main() {
// 	var wg sync.WaitGroup
// 	for _, filepath := range generateFilePath() {
// 		wg.Add(1)
// 		go processFile(filepath, &wg) // Pass by reference
// 	}

// 	wg.Wait() // Without this the main function will finish and won't give time for the routines to finish

// 	fmt.Printf("All files processed. Shutting down...")
// }

// 2. Real-time Web Request Counter
// Goal: Practice using sync/atomic for high-performance, thread-safe counters.
// Problem: Simulate a web server handling many requests concurrently. You need to keep an accurate count of the total requests received.
// Requirements:
// Create a global counter variable, totalRequests.
// Write a function handleRequest() that simulates handling a request. Inside this function, it should increment the totalRequests counter 100 times in a loop.
// In your main function, launch 1,000 goroutines, with each one calling handleRequest().
// Use a sync.WaitGroup to ensure the main function waits for all 1,000 goroutines to complete.
// After they are done, print the final value of totalRequests.
// Constraint:
// You must not use a sync.Mutex.
// The final printed value must be exactly 100,000.
// Primitive to use: sync/atomic package functions (e.g., atomic.AddUint64).

// Solution
// var totalRequests uint64

// func handleRequest(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for range 100 { // Equivalent to for i := range 100, 0 to 99
// 		atomic.AddUint64(&totalRequests, 1)
// 	}
// }

// func main() {
// 	var wg sync.WaitGroup
// 	for range 1000 {
// 		wg.Add(1)
// 		go handleRequest(&wg)
// 	}

// 	wg.Wait()

// 	fmt.Printf("%d\n", totalRequests)
// }

// 3. Task Queue Worker Pool
// Goal: Practice using channels to create a fixed-size worker pool for distributing jobs.
// Problem: You have 50 "jobs" (e.g., the numbers 1 to 50) that need to be processed, but you want to limit concurrency to only 4 active workers at any time.
// Requirements:
// Create two buffered channels:
// jobs := make(chan int, 50)
// results := make(chan int, 50)
// In main, send all 50 jobs to the jobs channel.
// Launch exactly 4 worker goroutines.
// Each worker should run in a loop, reading a job from the jobs channel.
// The worker should "process" the job (e.g., calculate its square: job * job), simulate work with time.Sleep(500 * time.Millisecond), and then send the result to the results channel.
// The main function must read all 50 results from the results channel and print them.
// Ensure your workers stop running once all jobs are complete (Hint: close the jobs channel).
// Primitive to use: channels

// Solution
// func worker(jobs <-chan int, results chan<- int) { // The channel types point if we're going to read or write
// 	for job := range jobs {
// 		time.Sleep(500 * time.Millisecond)
// 		results <- job * job
// 	}
// }

// func main() {
// 	// Create the channels
// 	jobs := make(chan int, 50)
// 	results := make(chan int, 50)

// 	// Create the jobs
// 	for i := 1; i <= 50; i++ {
// 		jobs <- i // Writing in the channel
// 	}
// 	// Close the jobs channel when it's done
// 	close(jobs)

// 	// Only 4 workers
// 	for i := 4; i <= 4; i++ {
// 		go worker(jobs, results) // Receives the current number, the jobs and the results channels
// 	}

// 	// Read results and print
// 	for i := 1; i <= 50; i++ {
// 		result := <-results
// 		fmt.Printf("Result: %d\n", result)
// 	}
// }

// 4. Service with Graceful Shutdown (Using OS Signals)
// Goal: Practice using channels, the select statement, and the os/signal package to handle system interrupts.
// Problem: You have a background "service" that runs in a goroutine, printing a "heartbeat" message every second. You need the application to catch the OS interrupt signal (e.g., Ctrl+C) and tell the service to stop cleanly before the program exits.
// Requirements:
// Use a sync.WaitGroup to allow main to wait for the service goroutine to finish.
// Create a "stop" channel (e.g., stopCh := make(chan struct{})) that will be used to signal the service to stop.
// Launch a service goroutine. Inside this goroutine:
// Call wg.Add(1) and defer wg.Done().
// Run an infinite for loop.
// Inside the loop, use a select statement to do one of two things:
// Wait for a 1-second tick (using time.NewTicker) and print "Service is running...".
// Wait for a signal on the stopCh. If received, print "Service stopping..." and return from the goroutine.
// In the main goroutine:
// Create a separate channel to receive OS signals (e.g., sigChan := make(chan os.Signal, 1)).
// Use signal.Notify to tell Go to send os.Interrupt signals to sigChan.
// Launch a new, anonymous goroutine whose only job is to block and wait for a signal on sigChan.
// When this signal-handling goroutine receives a signal, it should print "Caught signal, shutting down..." and then close(stopCh).
// After launching the service goroutine and the signal-handling goroutine, the main goroutine should call wg.Wait(). This will block main until the service has received the stop signal and has shut down (i.e., its wg.Done() is called).
// After wg.Wait() returns, main should print "Application shut down gracefully."
// Primitives/Packages to use: channels, select, sync.WaitGroup, os/signal, time

// Concurrency patterns training
// 1. Worker pools
// Distributes tasks across multiples goroutines
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		results <- job * 2 // simulate task process
	}
}
func main() {
	jobs, results := make(chan int, 100), make(chan int, 100)

}
