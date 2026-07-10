package main

import (
	"fmt"
	"jobprocessor/internal"
	"sync"
)


func main() {
	jobs, numOfJobs := work.CodeProvider()
	var wg sync.WaitGroup
	wg.Add(numOfJobs)

	for job := range jobs {
		res := ""
		go work.CodeResolver(&wg, job)
		fmt.Println(res)
	}
	
	wg.Wait()
}