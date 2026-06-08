package main

import (
	MTR "CFP/methodsTypesAndRoutines"
	"fmt"
	"sync"
)

func main() {
	fileLogs := []string{"Data.txt", "Network.txt", "Running.txt", "Clean-error.txt", "Append.txt", "Query-error.txt"}
	bufChan := make(chan MTR.LogResult, len(fileLogs)) // Helps strore certain amount of data, regardless if there's a receiver to collect or not.
	var wg sync.WaitGroup

	for _, fileName := range fileLogs {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			bufChan <- MTR.ProcessLog(fileName)
		}(&wg)
	}

	wg.Wait()
	close(bufChan)

	successes := make([]MTR.LogResult, 0)
	failures := make([]MTR.LogResult, 0)
	for val := range bufChan {
		if val.HasError {
			failures = append(failures, val)
			continue
		}
		successes = append(successes, val)

	}

	for _, success := range successes{
		name, numOfLine := success.FileName, success.LineCount
		fmt.Printf("✔  %s        — %d lines\n", name, numOfLine)
	}

	for _, failure := range failures{
		name, errMsg := failure.FileName, failure.ErrorMsg
		fmt.Printf("✗  %s        — %s lines\n", name, errMsg)
	}

	processed := len(successes) + len(failures)
	fmt.Printf("========================\nFiles Processed : %d\nSucceeded : %d\nFailed : %d\n========================\n", processed, len(successes), len(failures))
}
