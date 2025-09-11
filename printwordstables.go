package student

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, w := range a {
		for _, r := range w {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
