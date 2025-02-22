## Multithreading comparison Go vs Python

## Go
 go run (file path)
- Execution Time: 56.161084ms
- Uses all CPU cores effectively
- Lightweight goroutines → faster execution


## Python 
  Python3 (file path)
- Execution Time: 6.392580986022949
- Due to Python’s GIL, threads run sequentially (not in parallel).
- Same performance as a single-threaded program.


## Why is Python so Slow in Multi-threading?
🚫 Global Interpreter Lock (GIL):

Only one Python thread can execute at a time, even on multi-core CPUs.
Threads don’t run in parallel for CPU-bound tasks (they just take turns).

🚫 Thread Overhead:

Python threads have high context-switching costs.
The GIL locks/unlocks the interpreter, slowing down execution.

✅ Go wins because:

No GIL → Threads truly run in parallel across all CPU cores.
Lightweight goroutines → Less overhead compared to Python threads.

# 🚀 Conclusion
Go’s goroutines are much faster than Python threads because they allow true multi-threading without GIL interference.
Python threads are only useful for I/O-bound tasks (e.g., network requests, database queries).