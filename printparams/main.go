package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	program := args[1:]
	for _, arg := range program {
		for _, c := range arg {
			z01.PrintRune(c)
		}
		z01.PrintRune('\n')
	}
}
