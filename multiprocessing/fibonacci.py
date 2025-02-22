import multiprocessing
import time

# Compute Fibonacci (CPU-intensive)
def fibonacci(n):
    if n <= 1:
        return n
    return fibonacci(n - 1) + fibonacci(n - 2)

# Worker function
def worker(n):
    return fibonacci(n)

if __name__ == "__main__":
    num_workers = 4
    numbers = [35] * num_workers  # Compute Fibonacci(35) in parallel

    start_time = time.time()

    # Use multiprocessing pool
    with multiprocessing.Pool(num_workers) as pool:
        results = pool.map(worker, numbers)

    for res in results:
        print("Result:", res)

    elapsed_time = time.time() - start_time
    print("Execution Time:", elapsed_time)
