package student

import "github.com/01-edu/z01"

func DealAPackOfCards(deck []int) {
	for i := 0; i < 4; i++ {
		printString("Player ")
		printInt(i + 1)
		printString(": ")

		for j := 0; j < 13; j++ {
			index := i*13 + j
			printInt(deck[index])
			if j < 12 {
				printString(", ")
			}
		}
		z01.PrintRune('\n')
	}
}

// Εκτυπώνει έναν ακέραιο αριθμό με z01.PrintRune
func printInt(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	digits := []rune{}
	for n > 0 {
		digits = append([]rune{rune(n%10 + '0')}, digits...)
		n /= 10
	}

	for _, d := range digits {
		z01.PrintRune(d)
	}
}

// Εκτυπώνει string χαρακτήρα-χαρακτήρα
func printString(s string) {
	for _, ch := range s {
		z01.PrintRune(ch)
	}
}
