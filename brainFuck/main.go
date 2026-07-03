package main

import (
	"fmt"
	"os"
)

func main() {
	// a := 1
	// add(&a)
	// fmt.Println(a)
	brainFuck(os.Args[1])
}

func brainFuck(code string) {
	allBytes := [2048]byte{}
	reference := &allBytes[0]
	pointBytes := allBytes[1:]
	movementCount := map[rune]int{
		'>': 0,
		'<': 0,
	}
	position := 0
	for _, c := range code {
		switch c {
		case '>':
			movementCount[c]++
			position = movementCount[c]
			pointBytes[position] = *reference
		case '<':
			movementCount[c]--
			position = movementCount[c]
			pointBytes[position] = *reference
		case '+':
			pointBytes[position] = *reference
			pointBytes[position]++
		case '-':
			fmt.Print()
		case '.':
			fmt.Print()
		case '[':
			fmt.Print()
		case ']':
			fmt.Print()
		default:
			fmt.Println("foreign operator is present in the code string")
		}
	}
	fmt.Println(pointBytes)

}
