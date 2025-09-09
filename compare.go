package student

func Compare(a, b string) int {
	runesA := []rune(a)
	runesB := []rune(b)
	if a == b {
		return 0
	}
	n := len(runesA)
	if len(runesB) < n {
		n = len(runesB)
	}
	for i := 0; i < n; i++ {
		if runesA[i] < runesB[i] {
			return -1
		} else if runesA[i] > runesB[i] {
			return 1
		}
	}
	if len(runesA) < len(runesB) {
		return -1
	} else if len(runesA) > len(runesB) {
		return 1
	}
	return 0
}
