package internal

import "os"

func CodeProvider() (chan string, int) {
	codes := os.Args[1:]
	jobs := make(chan string, len(codes))

	for _, code := range codes {
		jobs <- code
	}
	close(jobs)

	return jobs, len(codes)
}