package student

func ShoppingSummaryCounter(str string) map[string]int {
	result := make(map[string]int)
	word := ""

	for _, r := range str {
		if r == ' ' {
			// αν έχουμε σχηματίσει λέξη, την αποθηκεύουμε
			if word != "" {
				result[word]++
				word = ""
			}
		} else {
			// προσθέτουμε χαρακτήρα στη λέξη
			word += string(r)
		}
	}

	// καταχωρούμε και την τελευταία λέξη (αν υπάρχει)
	if word != "" {
		result[word]++
	}

	return result
}
