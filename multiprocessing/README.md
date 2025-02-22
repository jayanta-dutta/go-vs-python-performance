## Multitprocessing comparison Go vs Python

## Go
 go run (file path)
- Execution Time: 54.248875ms

## Python
 Python3 (file path)
 - Execution Time: 1.755803108215332

 # ⚡ Why is Go Faster?
- Goroutines are lightweight → Go spawns thousands of goroutines efficiently, whereas Python’s processes have high overhead.

- Efficient memory sharing → Python multiprocessing creates separate processes, while Go shares memory.

- Better CPU utilization → Go effectively schedules goroutines across CPU cores, whereas Python’s multiprocessing has inter-process communication (IPC) overhead.

# Conclusion
- Go is ~3-5x faster than Python for CPU-bound parallel computations.
- If performance matters, Go is the clear winner for parallel workloads.
- Python still works well, but its multiprocessing has higher memory and IPC costs.