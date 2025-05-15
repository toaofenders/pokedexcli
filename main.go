package main

import(
	"fmt"
	"strings"
	)

func main() {
	fmt.Println("Hello, World!")
}

func 	cleanInput(text string) []string {
	// Trim leading and trailing whitespace
	text = strings.TrimSpace(text)
	// Split the text into words based on whitespace
	words := strings.Fields(text)
	// Convert all words to lowercase
	for i := range words {
		words[i] = strings.ToLower(words[i])
	}
	return words

}