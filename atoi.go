package student

func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	neg := false
	start := 0

	if s[0] == '-' {
		neg = true
		start = 1
	} else if s[0] == '+' {
		start = 1
	}

	result := 0
	for i := start; i < len(s); i++ {
		r := s[i]
		if r < '0' || r > '9' {
			return 0
		}
		result = result*10 + int(r-'0')
	}

	if neg {
		result = -result
	}

	return result
}
