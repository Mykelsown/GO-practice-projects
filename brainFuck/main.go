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
	
	jump := make([]int, len(code))
	stack := make([]int, 0)
	for i := 0; i < len(code); i++ {
		switch code[i] {
		case '[':
			stack = append(stack, i)
		case ']':
			open := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			jump[open] = i
			jump[i] = open
		}
	}

	var tape [2048]byte
	ptr := 0
	out := make([]byte, 0)

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