package student

func Fibonacci(index int) int {
	if index < 0 {
		return -1
	} else if index == 0 || index == 1 {
		return index
	}
	return Fibonacci(-1) + Fibonacci(-2)
}
