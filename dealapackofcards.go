package student

import "github.com/01-edu/z01"

func DealAPackOfCards(deck []int) {
	players := 4
	cardsPerPlayer := len(deck) / players

	for i := 0; i < players; i++ {
		playerHeader := "Player " + string('1'+i) + ": "
		for _, r := range playerHeader {
			z01.PrintRune(r)
		}
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
