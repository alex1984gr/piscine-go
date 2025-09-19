package student

import "github.com/01-edu/z01"

func DealAPackOfCards(deck []int) {
	// Κάθε παίκτης τυπώνεται με Player 1:, Player 2:, κλπ.
	for i := 0; i < 4; i++ {
		z01.PrintRune('P')
		z01.PrintRune('l')
		z01.PrintRune('a')
		z01.PrintRune('y')
		z01.PrintRune('e')
		z01.PrintRune('r')
		z01.PrintRune(' ')
		z01.PrintRune(rune(i + 1 + '0')) // αριθμός παίκτη
		z01.PrintRune(':')
		z01.PrintRune(' ')

		// 3 κάρτες ανά παίκτη
		for j := 0; j < 3; j++ {
			val := deck[i*3+j]
			if val >= 10 {
				z01.PrintRune(rune(1 + '0'))
				z01.PrintRune(rune(val%10 + '0'))
			} else {
				z01.PrintRune(rune(val + '0'))
			}
			if j < 2 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
