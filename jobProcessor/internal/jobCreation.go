package internal

import (
	"fmt"
	"os"
)

func JobCreator() {
	codes := os.Args[1:]
	jobs := make(chan string, len(codes))

	for _, code := range codes {
		jobs <- code
	}
	close(jobs)

	for job := range jobs{
		fmt.Println(job)
	}
}