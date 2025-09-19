package student

func printStr(s string) {
	for i := 0; i < len(s); i++ {
		PrintRune(rune(s[i])) // Χωρίς z01.
	}
}

func printNum(n int) {
	if n == 0 {
		PrintRune('0')
		return
	}

	if n < 0 {
		PrintRune('-')
		n = -n
	}

	var digits []rune
	for n > 0 {
		digits = append(digits, rune(n%10)+'0')
		n /= 10
	}

	for i := len(digits) - 1; i >= 0; i-- {
		PrintRune(digits[i])
	}
}
