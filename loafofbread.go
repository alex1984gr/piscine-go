package student

import "github.com/01-edu/z01"

func LoafOfBread(str string) string {
	if len(str) < 5 {
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
