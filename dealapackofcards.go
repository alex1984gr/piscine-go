package student

import "github.com/01-edu/z01"

func DealAPackOfCards(deck []int) {
	for i := 0; i < 4; i++ {
		player := "Player " + string('1'+i) + ": "
		for _, r := range player {
			z01.PrintRune(r)
		}
		for j := 0; j < 3; j++ {
			card := deck[i*3+j]
			if card >= 10 {
				z01.PrintRune(rune('0' + card/10))
				z01.PrintRune(rune('0' + card%10))
			} else {
				z01.PrintRune(rune('0' + card))
			}
			if j < 2 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
