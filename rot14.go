package student

func Rot14(s string) string {
	result := ""
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			result += string(((r-'a'+14)%26 + 'a'))
		} else if r >= 'A' && r <= 'Z' {
			result += string(((r-'A'+14)%26 + 'A'))
		} else {
			result += string(r)
		}
	}
	return result
}
