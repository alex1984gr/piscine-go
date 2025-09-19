package student

func Max(a []int) int {
	if len(a) == 0 {
		return 0
	}
	maxVarlue := a[0]
	for _, v := range a {
		if v > maxVarlue {
			maxVarlue = v
		}
	}
	return maxVarlue
}
