package student

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}

	var comb []int
	printCombs(n, 0, comb)
	z01.PrintRune('\n')
}

func printCombs(n, start int, comb []int) {
	if len(comb) == n {
		for _, d := range comb {
			z01.PrintRune(rune(d + '0'))
		}
		if !isLastComb(n, comb) {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
		return
	}

	for i := start; i <= 10-(n-len(comb)); i++ {
		printCombs(n, i+1, append(comb, i))
	}
}

func isLastComb(n int, comb []int) bool {
	for i := 0; i < n; i++ {
		if comb[i] != i+10-n {
			return false
		}
	}
	return true
}
