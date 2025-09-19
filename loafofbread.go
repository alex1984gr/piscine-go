package student

func LoafOfBread(str string) string {
	if len(str) < 5 {
		return "Invalid Output\n"
	}
	result := ""
	temp := ""
	count := 0
	skipNext := false
	for _, ch := range str {
		if skipNext {
			skipNext = false
			continue
		}
		if ch != ' ' {
			temp += string(ch)
			count++
		}
		if count == 5 {
			result += temp
			temp = ""
			count = 0
			skipNext = true
		}
	}
	if temp != "" {
		result += temp
	}

	return result + "\n"
}
