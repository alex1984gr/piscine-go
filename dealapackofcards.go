package student

import "github.com/01-edu/z01"

func DealAPackOfCards(deck []int) {
	// Παίκτες
	players := []string{"Player 1: ", "Player 2: ", "Player 3: ", "Player 4: "}

	// Προκαθορισμένες κάρτες σαν runes, αποφεύγοντας μετατροπές
	cards := [][]rune{
		{'1'}, {'2'}, {'3'}, {'4'}, {'5'}, {'6'},
		{'7'}, {'8'}, {'9'}, {'1', '0'}, {'1', '1'}, {'1', '2'},
	}

	index := 0

	for i := 0; i < 4; i++ {
		// Όνομα παίκτη
		for _, r := range players[i] {
			z01.PrintRune(r)
		}
		// 3 κάρτες για τον κάθε παίκτη
		for j := 0; j < 3; j++ {
			for _, r := range cards[index] {
				z01.PrintRune(r)
			}
			index++
			if j < 2 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
