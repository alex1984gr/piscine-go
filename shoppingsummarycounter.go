package student

func ShoppingSummaryCounter(str string) map[string]int {
	result := make(map[string]int)
	word := ""

	for _, r := range str {
		if r == ' ' {
			// μόνο αν υπάρχει λέξη, την μετράμε
			if word != "" {
				result[word]++
				word = ""
			}
		} else {
			word += string(r)
		}
	}

	// τελευταία λέξη αν έχει μείνει
	if word != "" {
		result[word]++
	}

	return result
}
