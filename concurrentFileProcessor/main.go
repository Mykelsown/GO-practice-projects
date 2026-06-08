package main

import (
	MTR "CFP/methodsTypesAndRoutines"
	"fmt"
)

func main() {
	fileLogs := []string{"Data.txt", "Network.txt", "Running.txt", "Clean-error.txt", "Append.txt", "Query-error.txt"}
	bufChan := make(chan string,  len(fileLogs))
	
	fmt.Println(MTR.ProcessLog("nERRORw.log"))
}