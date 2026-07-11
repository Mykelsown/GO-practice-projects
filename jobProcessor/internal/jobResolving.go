package work

import (
	"strings"
	"sync"
	"fmt"
)

// CodeResolver does the job i.e the workers: In this case the job is to convert a certain series of codes(operators), into a redadable alphabetical sentence(string).
func CodeResolver(job string, nWorkers int) chan string {
	finalizedChan := make(chan string, nWorkers)
	var wg sync.WaitGroup
	wg.Add(nWorkers)
	
	for i := 1; i <= nWorkers; i++ {
		go func(){
			fmt.Println(job, "other ", nWorkers)
			defer wg.Done()
			finalizedChan <- translate(job)
		}()
	} 

	go func() {
		wg.Wait()
		close(finalizedChan)
	}()
			
	return finalizedChan
}

// translate is a tool for the workers. It a conversion tool specified for this job received
func translate(codes string) string {
	// Build of the backtracking and forwarding logic; acheived by just swapping the position of the corresponding open and close with each other, so that they can be target in the logic that does the manipulation of the string
	positionSwapped := make([]int, len(codes))
	bracketsPositionCont := make([]int, 0)
	for i, code := range codes {
		switch code {
		case '[':
			bracketsPositionCont = append(bracketsPositionCont, i)
		case ']':
			openIndex := bracketsPositionCont[len(bracketsPositionCont)-1]
			bracketsPositionCont = bracketsPositionCont[:len(bracketsPositionCont)-1] // Removes the open bracket on finding the closing bracket
			positionSwapped[openIndex] = i                                            // i is te position of the current closed bracket, and it can only be a match with the last open, so it position is swapped with the correspondin open index
			positionSwapped[i] = openIndex
		}
	}

	// Logic that does the manipulation; convertion of the code string to a readable alphabetical string
	var(
		ptr [2048]byte
		movementCount int
		i int
		res strings.Builder
	)

	for i = 0; i < len(codes); i++{
		switch codes[i] {
		case '>':
			movementCount++
		case '<':
			movementCount--
		case '+':
			ptr[movementCount]++
		case '-':
			ptr[movementCount]--
		case '.':
			res.WriteByte(ptr[movementCount])
		case '[':
			if ptr[movementCount] == 0 {
				i = positionSwapped[i] //now note that position swapped will now be the position of the closed bracket.
			}
		case ']':
			if ptr[movementCount] != 0 {
				i = positionSwapped[i]
			}
		}
	} 

	return res.String()
}
