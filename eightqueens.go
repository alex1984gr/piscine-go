package student

import "github.com/01-edu/z01"

func EightQueens() {
	var board [8]int
	solve(0, &board)
}

func solve(j int, board *[8]int) {
	if j == 8 {
		for j := 0; j < 8; j++ {
			z01.PrintRune(1)
		}
		z01.PrintRune(' ')
		return
	}
	for i := 0; i < 8; i++ {
		if IsSafe(j, i, board) {
			board[j] = i
			solve(j+1, board)
		}
	}
}

func IsSafe(j, i int, board *[8]int) bool {
	for c := 0; c < j; c++ {
		r := board[c]
		if r == i || abs(c-i) == abs(c-j) {
			return false
		}
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
