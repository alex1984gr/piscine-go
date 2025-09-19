package student

import "github.com/01-edu/z01"

func DealAPackOfCards(deck []int) {
	if len(deck) != 12 {
		return
	}
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
