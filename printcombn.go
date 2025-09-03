package student

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}

	comb := make([]int, n)
	generateComb(0, 0, n, comb, true)
	z01.PrintRune('\n')
}

func generateComb(pos int, start int, n int, comb []int, first bool) {
	if pos == n {
		if !first {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
		for i := 0; i < n; i++ {
			z01.PrintRune(rune('0' + comb[i]))
		}
		return
	}

	for i := start; i <= 9-(n-pos); i++ {
		comb[pos] = i
		generateComb(pos+1, i+1, n, comb, first && pos == 0 && i == 0)
		first = false
	}
}
