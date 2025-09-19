package student

import "github.com/01-edu/z01"

func dealAPackOfCards(deck []int) {
	players := 4
	cardsPerPlayer := len(deck) / players
	for i := 0; i < players; i++ {
		z01.PrintRune('P')
		z01.PrintRune('l')
		z01.PrintRune('a')
		z01.PrintRune('y')
		z01.PrintRune('e')
		z01.PrintRune('r')
		z01.PrintRune(' ')
		z01.PrintRune(rune('1' + i))
		z01.PrintRune(':')
		z01.PrintRune(' ')
		for j := 0; j < cardsPerPlayer; j++ {
			card := deck[i*cardsPerPlayer+j]
			if card >= 10 {
				z01.PrintRune(rune('0' + card/10))
				z01.PrintRune(rune('0' + card%10))
			} else {
				z01.PrintRune(rune('0' + card))
			}
			if j != cardsPerPlayer-1 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}

}
