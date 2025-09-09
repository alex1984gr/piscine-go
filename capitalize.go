package student

func Capitalize(s string) string {
	result := ""
	newWord := true

	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			if newWord {
				if r >= 'a' && r <= 'z' {
					r = r - ('a' - 'A')
				}
				newWord = false
			} else {
				if r >= 'A' && r <= 'Z' {
					r = r + ('a' - 'A')
				}
			}
		} else {
			newWord = true
		}
		result += string(r)
	}
	return result
}
