package internal

import "os"

func CodeProvider() (<-chan string, int) {
	codes := os.Args[1:]
	jobs := make(chan string)
	
	for _, code := range codes {
		go func() {
			jobs <- code
			close(jobs)
		}()
	}

	return jobs, len(codes)
}