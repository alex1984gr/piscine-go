package student

func ShoppingSummaryCounter(str string) map[string]int {
	summary := map[string]int{}
	word := ""

	for _, ch := range str {
		if ch != ' ' {
			word += string(ch)
		} else if word != "" {
			summary[word]++
			word = ""
		}
		if word != "" {
			summary[word]++
		}
	}
	return summary
}
