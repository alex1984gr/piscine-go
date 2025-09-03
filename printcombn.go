package student

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}

	comb := make([]int, n)
	generateComb(0, 0, n, comb)
	z01.PrintRune('\n')
}

func generateComb(pos int, start int, n int, comb []int) {
	if pos == n {
		for i := 0; i < n; i++ {
			z01.PrintRune(rune('0' + comb[i]))
		}
		if !isLast(comb, n) {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
		return
	}
	for i := start; i <= 9-(n-pos); i++ {
		comb[pos] = i
		generateComb(pos+1, i+1, n, comb)
	}
}

func isLast(comb []int, n int) bool {
	for i := 0; i < n; i++ {
		if comb[i] != 10-n+i {
			return false
		}
	}
	return true
}
