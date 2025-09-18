package student

import "strings"

func ShoppingSummaryCounter(str string) map[string]int {
	summary := map[string]int{}
	words := strings.Split(str, " ")

	for _, word := range words {
		summary[word] = summary[word] + 1
	}

	return summary
}
