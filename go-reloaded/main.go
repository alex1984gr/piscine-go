package main

import "fmt"

func main() {
	fmt.Println("Enter text:") // μήνυμα προς χρήστη
	input := ReadInput()       // διαβάζει όλο το input

	tokens := tokenize(input)              // σπάει σε tokens
	afterTag := filterAfterTag(tokens)     // κρατά μόνο tokens μετά από tag
	quotes := extractQuotes(afterTag)      // κρατά μόνο tokens μέσα σε εισαγωγικά
	vowelWords := filterVowelWords(quotes) // φιλτράρει λέξεις

	// Εκτυπώνουμε το αποτέλεσμα για να χρησιμοποιήσουμε τη μεταβλητή
	fmt.Println("Words that start with vowels:")
	printOutput(vowelWords)
}
