package work

import (
	"os"
	"sync"
)

//CodeProvider creates job, and wait for a job to be solved before passing the next.
func CodeProvider() <-chan string {
	codes := os.Args[1:]
	var wg sync.WaitGroup
	wg.Add(len(codes))
	jobs := make(chan string, len(codes))
	
	for _, code := range codes {
		go func() {
			defer wg.Done()
			jobs <- code
		}()
	}
		
		wg.Wait()
		close(jobs)
		
	return jobs
}