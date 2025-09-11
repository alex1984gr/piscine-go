package student

func ConcatParams(args []string) string {
	result := ""
	for i := 0; i < len(args); i++ {
		s := args[i]
		result = result + s
		if i != len(args)-1 {
			result = result + "\n"
		}
	}
	return result
}
