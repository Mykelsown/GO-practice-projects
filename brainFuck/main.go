package main

import (
	"fmt"
	"os"
)

func main() {
	// a := []int{0, 1, 2, 3, 4, 5}
	// i := 0
	// count := 0
	// for i = 0; i < len(a); i++{
	// 	if a[i] == 2 && count < 5 {
	// 		i = 0
	// 		count++
	// 	}
	// 	fmt.Println(i)
	// }

	brainFuck(os.Args[1])
}

func brainFuck(code string) {
	allBytes := [2048]byte{}
	reference := allBytes[0]
	// pointBytes := allBytes[1:]
	// pointBytes[0] = &reference
	movementCount := 0
	out := ""
	
	var(
		ffCommand bool
		revCommand bool
	)

	openBracketPosition := make([]int, 0)
	openBracketPosition = append(openBracketPosition, 0)
	OBC := 0

	var i int
	for i = 0 ; i < len(code); i++ {
		// For the backtracking logic
		if revCommand {
			i = openBracketPosition[OBC]
			revCommand = false
			OBC--
			continue
		}

		// The condition helps satisfy the criteria for skipping all operators until it gets to ] operator
		if ffCommand && code[i] != ']' {
			continue
		} else if ffCommand && code[i]== ']' {
			ffCommand = false
			continue
		}
		
		switch code[i]{
		case '>':
			movementCount++
			pointBytes[movementCount] = &reference
			fmt.Println(*pointBytes[movementCount], "up")
		case '<':
			movementCount--
			pointBytes[movementCount] = &reference
			fmt.Println(movementCount, "down")
		case '+':
			*pointBytes[movementCount]++
			fmt.Println(reference, "a")
		case '-':
			
			*pointBytes[movementCount]--
			fmt.Println(reference, "m ", movementCount)
		case '.':
			out += string(reference)
			fmt.Println(out, "out")
			reference = 0
		case '[':
			if *pointBytes[movementCount] == 0{
				ffCommand = true
			}
			openBracketPosition = append(openBracketPosition, i+1)
			OBC++
			fmt.Println(openBracketPosition, "Obcccc")
		case ']':
			if *pointBytes[movementCount] != 0{
				// fmt.Println(openBracketPosition, "obc")
				revCommand = true
			}
		default:
			fmt.Println("foreign operator is present in the code string")
		}

		
	}
	// fmt.Println(*reference)
}
