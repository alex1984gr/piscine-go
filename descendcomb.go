package student

import "github.com/01-edu/z01"

func printNumber(n int) {
	z01.PrintRune(rune(n/10 + '0'))
	z01.PrintRune(rune(n%10 + '0'))
}

func DescendComb() {
	first := true
	for i := 99; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if !first {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			printNumber(i)
			z01.PrintRune(' ')
			printNumber(j)
			first = false
		}
	}
}
