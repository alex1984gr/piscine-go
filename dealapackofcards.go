package student

import (
	"fmt"

	"github.com/01-edu/z01"
)

func DealAPackOfCards(deck []int) {
	for i := 0; i < 4; i++ {
		// prefix μόνο με Printf (επιτρέπεται)
		fmt.Printf("Player %d: ", i+1)

		for j := 0; j < 3; j++ {
			val := deck[i*3+j]

			// χωρίς rune cast, explicit για κάθε περίπτωση
			if val == 10 {
				z01.PrintRune('1')
				z01.PrintRune('0')
			} else if val == 11 {
				z01.PrintRune('1')
				z01.PrintRune('1')
			} else if val == 12 {
				z01.PrintRune('1')
				z01.PrintRune('2')
			} else if val == 1 {
				z01.PrintRune('1')
			} else if val == 2 {
				z01.PrintRune('2')
			} else if val == 3 {
				z01.PrintRune('3')
			} else if val == 4 {
				z01.PrintRune('4')
			} else if val == 5 {
				z01.PrintRune('5')
			} else if val == 6 {
				z01.PrintRune('6')
			} else if val == 7 {
				z01.PrintRune('7')
			} else if val == 8 {
				z01.PrintRune('8')
			} else if val == 9 {
				z01.PrintRune('9')
			}

			if j < 2 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
