package student

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	digits := []int{}
	for n > 0 {
		digit := n % 10
		digits = append(digits, digit)
		n = n / 10
	}
	for i := 0; i < len(digits)-1; i++ {
		for j := 0; j < len(digits)-1-i; j++ {
			if digits[j] > digits[j+1] {
				digits[j], digits[j+1] = digits[j+1], digits[j]
			}
		}
	}
	for _, d := range digits {
		z01.PrintRune(rune(d + '0'))
	}
}
