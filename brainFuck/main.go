package main

import (
	"fmt"
	"os"
)

func main() {
	code := ""
	if len(os.Args) > 1 { 
		code = os.Args[1]
	}
	brainFuck(code)
}

func brainFuck(code string) {
	// Helps match the position of the correlating open and close bracket with each other, for easy mapping
	jump := make([]int, len(code))
	stack := make([]int, 0)
	for i := 0; i < len(code); i++ {
		switch code[i] {
		case '[':
			stack = append(stack, i) // Stores the position of the open bracket seen
		case ']':
			open := stack[len(stack)-1] // Gets the last element value in the stack i.e last open bracket 
			stack = stack[:len(stack)-1] // Update the element in the stack by removing the last element
			// Swapping of the open bracket and it correlating closing bracket index
			jump[open] = i
			jump[i] = open
		}
	}

	var tape [2048]byte
	ptr := 0
	out := make([]byte, 0)

	// Loop does the the movement of the pointer, printing out of the final character, and backtracking and fast forwading of commands
	for i := 0; i < len(code); i++ {
		switch code[i] {
		case '>':
			ptr++
		case '<':
			ptr--
		case '+':
			tape[ptr]++
		case '-':
			tape[ptr]--
		case '.':
			out = append(out, tape[ptr])
		case '[':
			if tape[ptr] == 0 {
				i = jump[i]
			}
		case ']':
			if tape[ptr] != 0 {
				i = jump[i]
			}
		}
	}
	fmt.Println(string(out))
}