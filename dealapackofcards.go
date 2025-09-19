package student

import "github.com/01-edu/z01"

func DealAPackOfCards(deck []int) {
	for player := 0; player < 4; player++ {
		z01.PrintRune('P')
		z01.PrintRune('l')
		z01.PrintRune('a')
		z01.PrintRune('y')
		z01.PrintRune('e')
		z01.PrintRune('r')
		z01.PrintRune(' ')
		z01.PrintRune(rune('1' + player))
		z01.PrintRune(':')
		z01.PrintRune(' ')
		for card := 0; card < 3; card++ {
			num := deck[player*3+card]
			if num >= 10 {
				z01.PrintRune(rune('0' + num/10))
			}
			z01.PrintRune(rune('0' + num%10))
			if card < 2 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
