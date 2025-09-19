package student

func ShoppingSummaryCounter(str string) map[string]int {
	result := make(map[string]int)
	word := ""

	for i := 0; i < len(str); i++ {
		c := str[i]

		if c == ' ' {
			if word != "" {
				result[word]++
				word = ""
			}
		} else {
			word += string(c)
		}
	}

	// το τελευταίο word αν υπάρχει
	if word != "" {
		result[word]++
	}

	return result
}
