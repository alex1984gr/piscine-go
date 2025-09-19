package student

import "github.com/01-edu/z01"

func LoafOfBread(str string) string {
	if len(str) < 5 {
		out := "Invalid Output\n"
		for _, r := range out {
			z01.PrintRune(r)
		}
		return out
	}
	result := ""
	word := ""
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
		word += string(ch)
		count++
		if count == 5 {
			result += word
			word = ""
			count = 0
			skipNext = true
		}
	}
	result += word
	result += "\n"
	for _, r := range result {
		z01.PrintRune(r)
	}

	return result
}
