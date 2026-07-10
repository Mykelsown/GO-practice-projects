package main

import (
	"jobprocessor/internal"
	"sync"
)


func main() {
	jobs, numOfJobs := internal.CodeProvider()
	var wg sync.WaitGroup
	wg.Add(numOfJobs)

	for job := range jobs {
		go internal.CodeResolver(&wg, job, jobs)
	}
	
	wg.Wait()
}