package student

import "github.com/01-edu/z01"

func DealAPackOfCards(deck []int) {
	players := []string{"Player 1: ", "Player 2: ", "Player 3: ", "Player 4: "}
	cardStrings := [][]rune{
		{'1'}, {'2'}, {'3'}, {'4'}, {'5'}, {'6'},
		{'7'}, {'8'}, {'9'}, {'1', '0'}, {'1', '1'}, {'1', '2'},
	}

	cardIndex := 0
	for i := 0; i < 4; i++ {
		// Όνομα παίκτη
		for _, r := range players[i] {
			z01.PrintRune(r)
		}
		// 3 κάρτες
		for j := 0; j < 3; j++ {
			for _, r := range cardStrings[cardIndex] {
				z01.PrintRune(r)
			}
			cardIndex++
			if j < 2 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
