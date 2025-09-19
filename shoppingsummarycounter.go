package student

func ShoppingSummaryCounter(str string) map[string]int {
	result := make(map[string]int)
	word := ""

	for i, r := range str {
		if r == ' ' {
			if word != "" {
				result[word]++
				word = ""
			}
		} else {
			word += string(r)
		}
		// Αν είμαστε στο τέλος της συμβολοσειράς
		if i == len(str)-1 && word != "" {
			result[word]++
		}
	}
	return result
}
