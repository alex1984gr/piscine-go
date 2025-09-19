package piscine

import "github.com/01-edu/z01"

func DealAPackOfCards(deck []int) {
	for i := 0; i < 4; i++ {
		printStr("Player ")
		printNum(i + 1)
		printStr(": ")

		for j := 0; j < 13; j++ {
			index := i*13 + j
			printNum(deck[index])
			if j < 12 {
				printStr(", ")
			}
		}
		z01.PrintRune('\n')
	}
}

func printStr(s string) {
	for _, ch := range s {
		z01.PrintRune(ch)
	}
}

func printNum(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	digits := [10]rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	var arr [10]int
	i := 0
	for n > 0 {
		arr[i] = n % 10
		n /= 10
		i++
	}
	for j := i - 1; j >= 0; j-- {
		z01.PrintRune(digits[arr[j]])
	}
}
