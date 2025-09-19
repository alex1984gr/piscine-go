package student

import (
	"fmt"

	"github.com/01-edu/z01"
)

func DealAPackOfCards(deck []int) {
	for i := 0; i < 4; i++ {
		// εκτυπώνουμε το prefix "Player X: "
		fmt.Printf("Player %d: ", i+1)

		// εκτυπώνουμε τις 3 κάρτες του παίκτη
		for j := 0; j < 3; j++ {
			val := deck[i*3+j]

			// ειδική περίπτωση για 10, 11, 12
			if val == 10 {
				z01.PrintRune('1')
				z01.PrintRune('0')
			} else if val == 11 {
				z01.PrintRune('1')
				z01.PrintRune('1')
			} else if val == 12 {
				z01.PrintRune('1')
				z01.PrintRune('2')
			} else {
				// μονοψήφιο
				z01.PrintRune('0' + rune(val))
			}

			// κόμμα και κενό, εκτός από την τελευταία κάρτα
			if j < 2 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}

		// newline στο τέλος κάθε παίκτη
		z01.PrintRune('\n')
	}
}
