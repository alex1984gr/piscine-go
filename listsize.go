package student

func ListSize(l *List) int {
	count := 0
	courrent := l.Head
	for courrent != nil {
		count++
		courrent = courrent.Next
	}
	return count
}
