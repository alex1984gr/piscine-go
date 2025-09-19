package student

import "strings"

func ShoppingSummaryCounter(str string) map[string]int {
	result := make(map[string]int)
	words := strings.Fields(str) // σπάει το string σε slice με βάση τα κενά

	for _, word := range words {
		result[word]++
	}

	return result
}
