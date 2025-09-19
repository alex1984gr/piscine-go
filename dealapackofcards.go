package student

import "fmt"

func DealAPackOfCards(deck []int) {
	for i := 0; i < 4; i++ {
		fmt.Printf("Player %d: ", i+1)
		for j := 0; j < 13; j++ {
			index := i*13 + j
			fmt.Print(deck[index])
			if j < 12 {
				fmt.Print(", ")
			}
		}
		fmt.Println()
	}
}
