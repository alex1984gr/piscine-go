package student

import "github.com/01-edu/z01"

func DealAPackOfCards(deck []int) {
	playerNames := [][]int{
		{'P', 'l', 'a', 'y', 'e', 'r', ' ', 1, ':', ' '},
		{'P', 'l', 'a', 'y', 'e', 'r', ' ', 2, ':', ' '},
		{'P', 'l', 'a', 'y', 'e', 'r', ' ', 3, ':', ' '},
		{'P', 'l', 'a', 'y', 'e', 'r', ' ', 4, ':', ' '},
	}

	for i := 0; i < 4; i++ {
		for _, c := range playerNames[i] {
			z01.PrintRune(c)
		}
		for j := 0; j < 3; j++ {
			val := deck[i*3+j]
			if val >= 10 {
				z01.PrintRune(1 + '0')
				z01.PrintRune(val%10 + '0')
			} else {
				z01.PrintRune(val + '0')
			}
			if j < 2 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
