package student

import "github.com/01-edu/z01"

func DealAPackOfCards(deck []int) {
	players := []string{"Player 1: ", "Player 2: ", "Player 3: ", "Player 4: "}
	deckIndex := 0

	for i := 0; i < 4; i++ {
		for _, r := range players[i] {
			z01.PrintRune(r)
		}
		for j := 0; j < 3; j++ {
			card := deck[deckIndex]
			deckIndex++
			if card < 10 {
				z01.PrintRune(rune('0' + card))
			} else if card == 10 {
				z01.PrintRune('1')
				z01.PrintRune('0')
			} else if card == 11 {
				z01.PrintRune('1')
				z01.PrintRune('1')
			} else if card == 12 {
				z01.PrintRune('1')
				z01.PrintRune('2')
			}
			if j < 2 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
