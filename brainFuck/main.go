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
	reference := &allBytes[0]
	pointBytes := allBytes[1:]
	movementCount := 0
	out := ""
	
	var(
		ffCommand bool
		revCommand bool
	)

	openBracketPosition := make([]int, 0)
	openBracketPosition = append(openBracketPosition, 0)
	OBC := 0

	i := 0
	for i = 0 ; i < len(code); i++ {
		// For the backtracking logic
		if revCommand {
			i = openBracketPosition[OBC] + 1
			revCommand = false
			OBC--
		}
			
		switch code[i]{
		case '>':
			movementCount++
			pointBytes[movementCount] = *reference
			fmt.Println(pointBytes[movementCount], "up")
		case '<':
			movementCount--
			pointBytes[movementCount] = *reference
			fmt.Println(pointBytes[movementCount], "down")
		case '+':
			*reference++
			println(*reference, "a")
		case '-':
			*reference--
			println(*reference, "m")
		case '.':
			out += string(*reference)
			*reference = 0
		case '[':
			if pointBytes[movementCount] == 0{
				ffCommand = true
				openBracketPosition = append(openBracketPosition, i)
				OBC++
			}
		case ']':
			if pointBytes[movementCount] != 0{
				revCommand = true
			}
		default:
			fmt.Println("foreign operator is present in the code string")
		}

		// The condition helps satisfy the criteria for skipping all operators until it gets to ] operator
		if ffCommand {
			continue
		} else if ffCommand && code[i]== ']' {
			ffCommand = false
			continue
		}
	}
	fmt.Println(*reference)
	fmt.Println(out)
}
