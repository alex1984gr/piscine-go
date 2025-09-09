package student

func TrimAtoi(s string) int {
	sign := 1
	started := false
	result := 0

	for _, ch := range s {
		if ch == '-' && !started {
			sign = -1
		}
		if ch >= '0' && ch <= '9' {
			started = true
			result = result*10 + int(ch-'0')
		}
	}
	return sign * result
}
