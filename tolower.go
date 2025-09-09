package student

func ToLower(s string) string {
	result := ""
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			r = r - ('A' - 'a')
		}
		result += string(r)
	}
	return result
}
