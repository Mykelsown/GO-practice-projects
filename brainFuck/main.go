package main

import (
	"fmt"
	// "os"
)

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	// b := &a[0]
	// *b = 10
	i := 0
	count := 0
	for i = 0; i < len(a); i++{
		if a[i] == 2 && count < 5 {
			i = 0
			count++
		}
		fmt.Println(i)
	}
	// fmt.Println(a)
	// brainFuck(os.Args[1])
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

	for _, c := range code {
		if ffCommand && revCommand {
			continue
		} else if ffCommand && c == ']' {
			ffCommand = false
		}
			
		switch c {
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
		case '[':
			if pointBytes[movementCount] == 0{
				ffCommand = true
			}
		case ']':
			if pointBytes[movementCount] == 0{
				revCommand = true
			}
		default:
			fmt.Println("foreign operator is present in the code string")
		}
	}
	fmt.Println(*reference)
	fmt.Println(out)
}
