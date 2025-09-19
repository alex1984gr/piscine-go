package student

func ShoppingSummaryCounter(str string) map[string]int {
	result := make(map[string]int)
	word := ""

	for _, r := range str {
		if r == ' ' {
			if word != "" {
				result[word]++
				word = ""
			}
		} else {
			word += string(r)
		}
	}
	if word != "" {
		result[word]++
	}

	return result
}
