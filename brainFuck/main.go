package main

import (
	"fmt"
	"os"
)

func main() {
	// a := []int{0, 1, 2, 3, 4, 5}
	// b := &a[0]
	// *b = 10
	// fmt.Println(a)
	brainFuck(os.Args[1])
}

func brainFuck(code string) {
	allBytes := [2048]byte{}
	reference := &allBytes[0]
	pointBytes := allBytes[1:]
	movementCount := 0
	for _, c := range code {
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
			fmt.Print()
		case '[':
			fmt.Print()
		case ']':
			fmt.Print()
		default:
			fmt.Println("foreign operator is present in the code string")
		}
	}
	fmt.Println(len(pointBytes))

}
