# Go Practice Question — Concurrent File Processor
 
**Difficulty:** Intermediate  
**Domain:** Server tooling / Concurrency
 
---
 
## Topics covered
 
| Topic | Where it applies |
|---|---|
| Goroutines | One goroutine launched per log file |
| WaitGroups | Track when all goroutines have finished |
| Buffered channels | Hold results without blocking workers |
| Closing a channel | Signal the collector that no more results are coming |
| Range over channel | Drain all results after close |
 
---
 
## Scenario
 
You are building a concurrent log processor for a server monitoring tool. The system receives a batch of log filenames, processes each file in parallel using goroutines, sends the results through a channel, and collects them in the main goroutine. Some logs contain errors — those must be separated from successful results before the final report is printed.
 
---
 
## Requirements
 
### 1 — LogResult struct
Define a `LogResult` struct with the following fields:
- `Filename string`
- `LineCount int`
- `HasError bool`
- `ErrorMsg string`
This is what each worker will produce and send into the channel.
 
---
 
### 2 — processLog function
Write a function `processLog(filename string) LogResult` that simulates processing a log file:
 
- Use `time.Sleep` with a random duration between **100ms–500ms** to simulate I/O work.
- If the filename contains the word `"error"`, return a `LogResult` with `HasError: true` and a descriptive `ErrorMsg`.
- Otherwise return a result with a random `LineCount` between 100–1000.
---
 
### 3 — Setup in main
In `main`:
 
- Define a slice of at least **6 log filenames** — at least 2 should contain `"error"` in the name.
- Create a **buffered channel** of type `LogResult` with a buffer size equal to the number of log files.
- Add a comment explaining why a buffered channel is appropriate here.
---
 
### 4 — Goroutines and WaitGroup
- Launch one **goroutine** per log file. Each goroutine calls `processLog` and sends the result into the channel.
- Use a `sync.WaitGroup` to track when all goroutines have finished.
- Once all are done, **close the channel**.
> Close the channel *after* `wg.Wait()`, not inside a goroutine — only the sender side should close, and only once all senders are done.
 
---
 
### 5 — Collect results
Use a `for range` loop over the channel to collect all results after it is closed. Separate them into two slices:
- `successes []LogResult`
- `failures []LogResult`
---
 
### 6 — Print the report
Print a final summary report:
- For each success: print the filename and line count.
- For each failure: print the filename and error message.
- End with a total count of files processed, succeeded, and failed.
---
 
## Expected output (order may vary)
 
```
Processing 6 log files concurrently...
 
✔  server-2024.log        — 743 lines
✔  access-jan.log         — 218 lines
✔  metrics-daily.log      — 501 lines
✔  traffic-audit.log      — 894 lines
✗  error-db-timeout.log   — failed to parse error-db-timeout.log: critical error detected
✗  error-auth-service.log — failed to parse error-auth-service.log: critical error detected
 
══════════════════════════
  Files processed : 6
  Succeeded       : 4
  Failed          : 2
══════════════════════════
```
 
> The output order of results is non-deterministic — goroutines finish in whatever order the scheduler runs them. Your output will differ from the example above on every run. That is expected and correct behaviour.
 
---

## Possible extensions
 
- Add a `workerPool` pattern: instead of one goroutine per file, limit to N workers using a semaphore channel.
- Add a context with timeout — cancel all goroutines if total processing exceeds 2 seconds.
- Write results to a file using `os.WriteFile` instead of printing to stdout.
 