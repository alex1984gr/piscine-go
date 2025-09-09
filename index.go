package student

func Index(s string, toFind string) int {
	runesA := []rune(s)
	runesB := []rune(toFind)
	for i := 0; i <= len(runesA)-len(runesB); i++ {
		match := true
		for j := 0; j < len(runesB); j++ {
			if runesA[i+j] != runesB[j] {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}
	return -1
}
