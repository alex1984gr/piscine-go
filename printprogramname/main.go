package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	program := args[0]
	lastSlash := -1
	for i, c := range program {
		if c == '/' {
			lastSlash = i
		}
	}
	name := program[lastSlash+1:]
	for _, c := range name {
		z01.PrintRune(rune(c))
	}
	z01.PrintRune('\n')
}
