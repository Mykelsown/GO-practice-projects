package main

import (
	"fmt"
	work "jobprocessor/internal"
)


func main() {
	jobs := work.CodeProvider()
	
	var results <-chan string
	results = work.CodeResolver(jobs, 2)

	for result := range results {
		fmt.Print(result)
	}
}