package internal

import (
	"fmt"
	"strings"
	"sync"
)

func CodeResolver(wg *sync.WaitGroup) {
	defer wg.Done()

	jobs, numOfJobs := CodeProvider()
	finalizedChan := make(chan string, numOfJobs)

	for job := range jobs {
		finalizedChan <- translate(job)
	}
	close(finalizedChan)

	for finalized := range finalizedChan{
		fmt.Println(finalized)
	}
}

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
