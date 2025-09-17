package student

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) <= 1 {
		return true
	}

	increasing := true
	decreasing := true

	for i := 0; i < len(a)-1; i++ {
		cmp := f(a[i], a[i+1])
		if cmp > 0 {
			increasing = false
		}
		if cmp < 0 {
			decreasing = false
		}
	}

	return increasing || decreasing
}
