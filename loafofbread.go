package student

import "github.com/01-edu/z01"

func LoafOfBread(str string) string {
	if len(str) < 5 {
		return "Invalid Output\n"
	}

	count := 0
	skipNext := false

	for _, ch := range str {
		if skipNext {
			skipNext = false
			continue
		}

		if ch != ' ' {
			z01.PrintRune(ch)
			count++
		}

		if count == 5 {
			count = 0
			skipNext = true
		}
	}

	z01.PrintRune('\n')
	return ""
}
