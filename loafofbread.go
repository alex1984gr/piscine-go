package student

import "github.com/01-edu/z01"

func LoafOfBread(str string) string {
	if len(str) < 5 {
		z01.PrintRune('I')
		z01.PrintRune('n')
		z01.PrintRune('v')
		z01.PrintRune('a')
		z01.PrintRune('l')
		z01.PrintRune('i')
		z01.PrintRune('d')
		z01.PrintRune(' ')
		z01.PrintRune('O')
		z01.PrintRune('u')
		z01.PrintRune('t')
		z01.PrintRune('p')
		z01.PrintRune('u')
		z01.PrintRune('t')
		z01.PrintRune('\n')
		return "Invalid Output\n"
	}

	result := ""
	count := 0
	skipNext := false

	for _, ch := range str {
		if ch == ' ' {
			result += string(ch)
			continue
		}
		if skipNext {
			skipNext = false
			continue
		}
		result += string(ch)
		count++
		if count == 5 {
			count = 0
			skipNext = true
		}
	}

	result += "\n"

	for _, r := range result {
		z01.PrintRune(r)
	}

	return result
}
