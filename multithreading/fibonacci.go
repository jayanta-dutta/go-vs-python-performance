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
func worker(id int, n int, wg *sync.WaitGroup, results chan<- int) {
	defer wg.Done()
	fmt.Printf("Worker %d computing Fibonacci(%d)\n", id, n)
	results <- fibonacci(n)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // Use all CPU cores

	numThreads := 4
	n := 35
	var wg sync.WaitGroup
	results := make(chan int, numThreads)

	start := time.Now()

	// Start 4 goroutines (threads)
	for i := 1; i <= numThreads; i++ {
		wg.Add(1)
		go worker(i, n, &wg, results)
	}

	wg.Wait()
	close(results)

	// Collect results
	for res := range results {
		fmt.Println("Result:", res)
	}

	elapsed := time.Since(start)
	fmt.Println("Execution Time:", elapsed)
}
