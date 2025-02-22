import threading
import time

# Compute Fibonacci (CPU-intensive)
def fibonacci(n):
    if n <= 1:
        return n
    return fibonacci(n - 1) + fibonacci(n - 2)

# Worker function
def worker(n, results, thread_id):
    print(f"Thread {thread_id} computing Fibonacci({n})")
    results.append(fibonacci(n))

if __name__ == "__main__":
    num_threads = 4
    n = 35
    results = []
    threads = []

    start_time = time.time()

    # Start 4 threads
    for i in range(num_threads):
        thread = threading.Thread(target=worker, args=(n, results, i))
        threads.append(thread)
        thread.start()

    # Wait for all threads to complete
    for thread in threads:
        thread.join()

    for res in results:
        print("Result:", res)

    elapsed_time = time.time() - start_time
    print("Execution Time:", elapsed_time)
