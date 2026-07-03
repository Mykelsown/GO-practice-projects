package main

import "fmt"

func main() {
	// a := 1
	// add(&a)
	// fmt.Println(a)
	brainFuck()
}

func add(a *int) *int {
	z := 5
	y := &z
	*a = *a * *y
	return a
}

func brainFuck(code string) {
	allBytes := [2048]byte{}
	reference := &allBytes[0]
	pointBytes := allBytes[1:]
	movementCount := map[rune]int{
		'>': 0,
		'<': 0,
	}
	for _, c := range code {

		switch c {
		case '>':
			movementCount[c]++
			pointBytes[movementCount[c]] = *reference
		case '<':
			movementCount[c]--
			pointBytes[movementCount[c]] = *reference
		case '+':
			fmt.Print()
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
