package student

func ShoppingSummaryCounter(str string) map[string]int {
	result := make(map[string]int)
	word := ""

	for _, r := range str {
		if r == ' ' {
			// καταχωρούμε πάντα το τρέχον word, ακόμη κι αν είναι άδειο
			result[word]++
			word = ""
		} else {
			word += string(r)
		}
	}

	// καταχωρούμε και το τελευταίο πεδίο (ίσως κενό)
	result[word]++

	return result
}
