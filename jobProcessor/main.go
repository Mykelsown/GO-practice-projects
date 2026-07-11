package main

import (
	"fmt"
	"jobprocessor/internal"
)


func main() {
	jobs, numOfJobs := work.CodeProvider()
	
	var results chan string
	for job := range jobs {
		fmt.Println(job, "main")
		results = work.CodeResolver(job, numOfJobs)
	}

	for result := range results {
		fmt.Println(result)
	}
	
}