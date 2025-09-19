package student

func LoafOfBread(str string) string {
	myStr := ""
	word := ""
	skip := false
	for _, char := range str {
		if skip {
			skip = false
			continue
		}
		if char == ' ' {
			continue
		}
		word += string(char)
		if len(word) == 5 {
			myStr += word + " "
			word = ""
			skip = true
		}
	}
	if word != "" {
		myStr += word
	}
	if myStr == "" {
		return "\n"
	}
	if myStr[len(myStr)-1] == ' ' {
		myStr = myStr[:len(myStr)-1]
	}
	return myStr + "\n"
}
