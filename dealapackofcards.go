package student

import "github.com/01-edu/z01"

func DealAPackOfCards(deck []int) {
	players := []string{"Player 1: ", "Player 2: ", "Player 3: ", "Player 4: "}
	cardChars := [][]rune{
		{'1'}, {'2'}, {'3'}, {'4'}, {'5'}, {'6'},
		{'7'}, {'8'}, {'9'}, {'1', '0'}, {'1', '1'}, {'1', '2'},
	}

	for i := 0; i < 4; i++ {
		for _, r := range players[i] {
			z01.PrintRune(r)
		}
		for j := 0; j < 3; j++ {
			cardIndex := i*3 + j
			if j > 0 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			for _, r := range cardChars[deck[cardIndex]-1] {
				z01.PrintRune(r)
			}
		}
		z01.PrintRune('\n')
	}
}
