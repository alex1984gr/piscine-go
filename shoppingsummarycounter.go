package student

import "strings"

func ShoppingSummaryCounter(str string) map[string]int {
	result := make(map[string]int)
	items := strings.Split(str, " ")

	for _, item := range items {
		result[item]++
	}

	return result
}
