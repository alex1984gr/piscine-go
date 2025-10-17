package main

func main() {
	input := raedInput()
	tokens := tokenize(input)
	afterTag := FilterAfterTag(tokens)
	InsideQuotes := ExtractQuotes(afterTag)
	vowelWords := FilterVowelWords(InsideQuotes)
	PrintOutput(vowelWords)
}
