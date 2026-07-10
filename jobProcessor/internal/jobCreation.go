package work

import (
	"os"
	"sync"
)

func CodeProvider() (<-chan string, int) {
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
		
	return jobs, len(codes)
}