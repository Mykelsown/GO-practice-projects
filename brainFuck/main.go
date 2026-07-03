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
	pointBytes := [2048]byte{}
	reference := 
	movementCount := map[rune]int{
		'>': 0,
		'<': 0,
	}
	for _, c := range code {

		switch c {
		case '>':
			ind := movementCount[c]+1
			pointBytes[ind] = 
		case '<':
			
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
