package main

import (
	"bufio"   // Πακέτο για buffered I/O
	"fmt"     // Πακέτο για εκτύπωση
	"os"      // Πακέτο για ανάγνωση stdin
	"strings" // Πακέτο για επεξεργασία strings
)

// Διαβάζει όλο το input από stdin
func ReadInput() []rune {
	scanner := bufio.NewScanner(os.Stdin) // δημιουργεί ένα σκάνερ που διαβάζει γραμμές από το τερματικό
	text := ""                            // δημιουργεί μεταβλητή text τύπου string για αποθήκευση κειμένου
	for scanner.Scan() {                  // loop που τρέχει όσο υπάρχουν γραμμές για να διαβάσει ο scanner
		text += scanner.Text() + " " // προσθέτει τη γραμμή που διάβασε το scanner στο string text
	}
	return []rune(text) // μετατρέπει όλο το string σε slice από runes
}

// Σπάει το κείμενο σε tokens: <tags>, "quotes", words
func tokenize(input []rune) []string {
	tokens := []string{} // slice για να αποθηκεύσει τα tokens
	current := ""        // προσωρινή μεταβλητή για την τρέχουσα λέξη

	for _, ch := range input { // loop σε κάθε χαρακτήρα του input
		if ch == '<' || ch == '>' || ch == '"' || ch == ' ' { // αν είναι ειδικός χαρακτήρας
			if current != "" { // αν υπάρχει κάτι στο current
				tokens = append(tokens, current) // προσθέτει την τρέχουσα λέξη στα tokens
				current = ""                     // καθαρίζει το current
			}
			if ch != ' ' { // αν δεν είναι κενό
				tokens = append(tokens, string(ch)) // προσθέτει τον ειδικό χαρακτήρα στα tokens
			}
		} else {
			current += string(ch) // αλλιώς προσθέτει τον χαρακτήρα στο current
		}
	}

	if current != "" { // αν υπάρχει υπόλοιπο στο current
		tokens = append(tokens, current) // το προσθέτει στα tokens
	}

	return tokens // επιστρέφει όλα τα tokens
}

// Κρατά tokens που είναι μετά από tags
func filterAfterTag(tokens []string) []string {
	result := []string{} // slice για τα tokens μετά από tag
	insideTag := false   // κατάσταση αν είμαστε μετά από > ή όχι

	for i, tok := range tokens {
		if tok == ">" { // αν βρούμε >
			insideTag = true // μπαίνουμε σε κατάσταση "μετά από tag"
			continue
		}
		if tok == "<" && i+1 < len(tokens) && tokens[i+1] != "/" { // αν βρούμε νέο < (εκτός από κλείσιμο)
			insideTag = false // βγαίνουμε από κατάσταση "μετά από tag"
		}
		if insideTag { // αν είμαστε μετά από tag
			result = append(result, tok) // προσθέτουμε το token στο αποτέλεσμα
		}
	}

	return result // επιστρέφει τα tokens μετά από tags
}

// Παίρνει μόνο το περιεχόμενο μέσα στα εισαγωγικά
func extractQuotes(tokens []string) []string {
	result := []string{}  // slice για λέξεις μέσα σε ""
	insideQuotes := false // κατάσταση αν είμαστε μέσα σε εισαγωγικά

	for _, tok := range tokens {
		if tok == `"` { // αν βρούμε "
			insideQuotes = !insideQuotes // αλλάζουμε κατάσταση
			continue
		}
		if insideQuotes { // αν είμαστε μέσα σε εισαγωγικά
			result = append(result, tok) // προσθέτουμε το token στο αποτέλεσμα
		}
	}

	return result // επιστρέφει τα tokens μέσα σε εισαγωγικά
}

// Επιστρέφει true αν το word ξεκινά με φωνήεν
func startsWithVowel(word string) bool {
	if len(word) == 0 { // αν η λέξη είναι κενή
		return false // επιστρέφει false
	}
	vowels := "aeiouAEIOU"                             // όλα τα φωνήεντα
	return strings.ContainsRune(vowels, rune(word[0])) // ελέγχει αν το πρώτο γράμμα είναι φωνήεν
}

// Φιλτράρει λέξεις που ξεκινούν με φωνήεν
func filterVowelWords(words []string) []string {
	result := []string{} // slice για λέξεις που ξεκινούν με φωνήεν
	for _, word := range words {
		if startsWithVowel(word) { // αν ξεκινά με φωνήεν
			result = append(result, word) // προσθέτει τη λέξη στο αποτέλεσμα
		}
	}
	return result // επιστρέφει το αποτέλεσμα
}

// Εκτυπώνει το τελικό αποτέλεσμα
func printOutput(words []string) {
	for _, w := range words { // για κάθε λέξη
		fmt.Print(w, " ") // εκτυπώνει με κενό
	}
	fmt.Println() // προσθέτει αλλαγή γραμμής στο τέλος
}
