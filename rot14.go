package student

import "github.com/01-edu/z01"

func Rot14(s string) {
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			z01.PrintRune(((r-'a'+14)%26 + 'a'))
		} else if r >= 'A' && r <= 'Z' {
			z01.PrintRune(((r-'A'+14)%26 + 'A'))
		} else {
			z01.PrintRune(r)
		}
	}
	z01.PrintRune('\n')
}
