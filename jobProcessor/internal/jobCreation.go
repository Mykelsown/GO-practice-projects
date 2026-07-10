package internal

import "os"

func JobCreator() chan string {
	codes := os.Args[1:]
	jobs := make(chan string, len(codes))

	for _, code := range codes {
		jobs <- code
	}
	close(jobs)

	return jobs
}