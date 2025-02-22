package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Compute Fibonacci (CPU-intensive)
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// Worker function
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range jobs {
		results <- fibonacci(n)
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // Use all CPU cores

	numJobs := 4
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	var wg sync.WaitGroup

	start := time.Now()

	// Start 4 workers
	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// Send jobs
	for i := 0; i < numJobs; i++ {
		jobs <- 35 // Compute Fibonacci(35) in parallel
	}
	close(jobs)

	// Wait for workers to finish
	wg.Wait()
	close(results)

	// Collect results
	for res := range results {
		fmt.Println("Result:", res)
	}

	elapsed := time.Since(start)
	fmt.Println("Execution Time:", elapsed)
}
