package student

import "github.com/01-edu/z01"

func DealAPackOfCards(deck [12]int) {
	players := 4
	cardsPerPlayer := 3
	index := 0
	for p := 1; p <= players; p++ {
		z01.PrintRune('P')
		z01.PrintRune('l')
		z01.PrintRune('a')
		z01.PrintRune('y')
		z01.PrintRune('e')
		z01.PrintRune('r')
		z01.PrintRune(' ')
		z01.PrintRune(byte('0' + p))
		z01.PrintRune(':')
		z01.PrintRune(' ')

		for c := 0; c < cardsPerPlayer; c++ {
			card := deck[index]
			if card < 10 {
				z01.PrintRune(byte('0' + card))
			} else {
				if card == 10 {
					z01.PrintRune('1')
					z01.PrintRune('0')
				} else if card == 11 {
					z01.PrintRune('1')
					z01.PrintRune('1')
				} else if card == 12 {
					z01.PrintRune('1')
					z01.PrintRune('2')
				}
			}
			index++
			if c < cardsPerPlayer-1 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
