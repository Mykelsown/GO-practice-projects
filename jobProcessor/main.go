package main

import (
	"jobprocessor/internal"
	"sync"
)


func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go internal.JobResolver(&wg)
	wg.Wait()
}